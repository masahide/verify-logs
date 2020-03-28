package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/kelseyhightower/envconfig"
)

type param struct {
	LogGroupName   string        `default:"/aws/ecs/verify-logs"`
	Pattern        string        `default:""`
	DummySize      int           `default:"1000"`
	DummyStringLen int           `default:"300"`
	MaxLines       int           `default:"1000000"`
	SetOfLines     int           `default:"2000"`
	Wait           time.Duration `default:"5s"`
}

func getDummyData(p param) []string {
	s := make([]string, p.DummySize)
	for i := 0; i < p.DummySize; i++ {
		s[i] = strings.Repeat("a", p.DummyStringLen)
	}
	return s
}

func notIn(ss []string, s string) bool {
	for i := range ss {
		if s == ss[i] {
			return false
		}
	}
	return true
}

func tail(result chan<- string, p param) {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	svc := cloudwatchlogs.New(sess, aws.NewConfig())
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
	No   int    `json:"log_number"`
	Data string `json:"data"`
}

func outputLogs(data []string, p param) {
	for i := 0; i < p.MaxLines; i++ {
		l := logData{No: i, Data: data[i%len(data)]}
		j, err := json.Marshal(l)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", j)
		if i%p.SetOfLines == 0 {
			time.Sleep(p.Wait)
		}
	}
}

func check(data []string, p param) {
	res := make(chan string, p.SetOfLines*10)
	go tail(res, p)
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
		if data.No == index {
			index++
		} else {
			log.Fatalf("err: index(%d) != data.No(%d), line:%s", index, data.No, r)
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
	data := getDummyData(p)
	go outputLogs(data, p)
	check(data, p)
}
