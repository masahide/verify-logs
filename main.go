package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/docker/docker/daemon/logger"
	"github.com/docker/docker/daemon/logger/awslogs"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type param struct {
	LogGroupName    string        `default:"test-verify-logs"`
	Regeon          string        `envconfig:"AWS_REGION" default:"ap-northeast-1"`
	LogStreamPrefix string        `default:"verify-logs"`
	Pattern         string        `default:""`
	DummyStringLen  int           `default:"30"`
	MaxLines        int           `default:"1000"`
	SetOfLines      int           `default:"5"`
	Wait            time.Duration `default:"3s"`
}

func createLogGroup(svc cloudwatchlogsiface.CloudWatchLogsAPI, p param) {
	_, err := svc.CreateLogGroup(&cloudwatchlogs.CreateLogGroupInput{LogGroupName: &p.LogGroupName})
	if err != nil {
		switch err.(type) {
		case *cloudwatchlogs.ResourceAlreadyExistsException:
			log.Printf("already %s", p.LogGroupName)
			return
		}
		log.Fatal(err)
	}
	log.Printf("wait...")
	time.Sleep(5 * time.Second)
}

func notIn(ss []string, s string) bool {
	for i := range ss {
		if s == ss[i] {
			return false
		}
	}
	return true
}

// tail reproduces the implementation of aws-cliv2 logs tail command.
// see: https://github.com/aws/aws-cli/blob/v2/awscli/customizations/logs/tail.py#L298-L313
func tail(svc cloudwatchlogsiface.CloudWatchLogsAPI, result chan<- string, p param) {
	var nextToken *string
	startTime := aws.Int64(time.Now().Unix() * 1000)
	lastEventIDs := []string{}
	for {
		input := &cloudwatchlogs.FilterLogEventsInput{
			StartTime:     startTime,
			LogGroupName:  aws.String(p.LogGroupName),
			NextToken:     nextToken,
			FilterPattern: aws.String(p.Pattern),
		}
		output, err := svc.FilterLogEvents(input)
		if err != nil {
			log.Fatal(err)
		}
		ids := make([]string, len(output.Events))
		for i, event := range output.Events {
			if notIn(lastEventIDs, aws.StringValue(event.EventId)) {
				result <- aws.StringValue(event.Message)
			}
			ids[i] = aws.StringValue(event.EventId)
		}
		lastEventIDs = ids
		if output.NextToken != nil {
			nextToken = output.NextToken
		}
		if nextToken != nil {
			time.Sleep(2 * time.Second)
		}
	}
}

type logData struct {
	Time time.Time `json:"time"`
	No   int       `json:"log_number"`
	Data string    `json:"data"`
}

func newLogger(p param) logger.Logger {
	info := logger.Info{Config: map[string]string{
		"awslogs-group":         p.LogGroupName,
		"awslogs-region":        p.Regeon,
		"awslogs-stream-prefix": p.LogStreamPrefix,
	},
		ContainerID:   uuid.Must(uuid.NewRandom()).String(),
		ContainerName: p.LogGroupName,
	}
	alog, err := awslogs.New(info)
	if err != nil {
		log.Fatal(err)
	}
	return alog
}

// outputLogs  Log output to cloudwatchlogs using `docker awslogs driver`
func outputLogs(p param) {
	alog := newLogger(p)
	for i := 0; i < p.MaxLines; i++ {
		l := logData{Time: time.Now(), No: i, Data: strings.Repeat("a", p.DummyStringLen)}
		j, err := json.Marshal(l)
		if err != nil {
			log.Fatal(err)
		}
		alog.Log(&logger.Message{Line: j, Timestamp: time.Now()})
		if i%p.SetOfLines == 0 {
			alog.Close()
			time.Sleep(p.Wait)
			alog = newLogger(p)
		}
	}
}

func check(svc cloudwatchlogsiface.CloudWatchLogsAPI, p param) {
	res := make(chan string, p.SetOfLines*10)
	go tail(svc, res, p)
	index := 0
	for {
		r := <-res
		// find json string
		head := strings.Index(r, "{")
		if head == -1 {
			continue
		}

		// parse JSON log
		var data logData
		err := json.Unmarshal([]byte(r[head:]), &data)
		if err != nil {
			log.Printf("faild json.Unmarshal:%s", err)
			continue
		}
		log.Printf("%s\n", r)

		if data.No == index {
			index++
		} else {
			log.Printf("err: index(%d) != data.No(%d), line:%s", index, data.No, r)
		}
		if index >= p.MaxLines {
			log.Printf("Everything is fine!")
			return
		}
	}
}

func main() {
	var p param
	err := envconfig.Process("", &p)
	if err != nil {
		log.Fatal(err)
	}
	svc := cloudwatchlogs.New(session.Must(session.NewSession()), aws.NewConfig())
	createLogGroup(svc, p)
	go outputLogs(p)
	check(svc, p)
}
