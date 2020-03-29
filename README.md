# verify-logs



[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/masahide/verify-logs)](https://hub.docker.com/repository/docker/masahide/verify-logs)[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/masahide/verify-logs)](https://hub.docker.com/r/masahide/verify-logs/builds)


## Usage

```bash
# use ~/.aws/config 
docker run -e AWS_PROFILE=profile-name -e AWS_SDK_LOAD_CONFIG=1 -e AWS_DEFAULT_REGION=ap-northeast-1 -v $HOME:/root masahide/verify-logs

# set loggroup-name
docker run -e LogGroupName=loggroup-name -e AWS_PROFILE=profile-name -e AWS_SDK_LOAD_CONFIG=1 -e AWS_DEFAULT_REGION=ap-northeast-1 -v $HOME:/root masahide/verify-logs
```


## Execution example of failure reproduction
```bash
23:47 $ docker run -e AWS_PROFILE=profile_name -e AWS_SDK_LOAD_CONFIG=1 -e AWS_DEFAULT_REGION=ap-northeast-1 -v $HOME:/root masahide/verify-logs
2020/03/29 14:47:09 wait...
2020/03/29 14:47:15 {"time":"2020-03-29T14:47:14.7961866Z","log_number":0,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:19 {"time":"2020-03-29T14:47:18.0517547Z","log_number":1,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:19 {"time":"2020-03-29T14:47:18.0517942Z","log_number":2,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:19 {"time":"2020-03-29T14:47:18.0518217Z","log_number":3,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:19 {"time":"2020-03-29T14:47:18.0518531Z","log_number":4,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:19 {"time":"2020-03-29T14:47:18.0518847Z","log_number":5,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:23 {"time":"2020-03-29T14:47:21.3043291Z","log_number":6,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:23 {"time":"2020-03-29T14:47:21.3043721Z","log_number":7,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:23 {"time":"2020-03-29T14:47:21.3044033Z","log_number":8,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:23 {"time":"2020-03-29T14:47:21.3044321Z","log_number":9,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:23 {"time":"2020-03-29T14:47:21.3044631Z","log_number":10,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:27 {"time":"2020-03-29T14:47:24.5748432Z","log_number":11,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:27 {"time":"2020-03-29T14:47:24.5748897Z","log_number":12,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:27 {"time":"2020-03-29T14:47:24.57492Z","log_number":13,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:27 {"time":"2020-03-29T14:47:24.5749488Z","log_number":14,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:27 {"time":"2020-03-29T14:47:24.5749794Z","log_number":15,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:30 {"time":"2020-03-29T14:47:27.8163598Z","log_number":16,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:30 {"time":"2020-03-29T14:47:27.8164119Z","log_number":17,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:30 {"time":"2020-03-29T14:47:27.8164428Z","log_number":18,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:30 {"time":"2020-03-29T14:47:27.8164787Z","log_number":19,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:30 {"time":"2020-03-29T14:47:27.8165096Z","log_number":20,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:34 {"time":"2020-03-29T14:47:31.0675533Z","log_number":21,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:34 {"time":"2020-03-29T14:47:31.0676005Z","log_number":22,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:34 {"time":"2020-03-29T14:47:31.0676341Z","log_number":23,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:34 {"time":"2020-03-29T14:47:31.0676733Z","log_number":24,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:34 {"time":"2020-03-29T14:47:31.0677046Z","log_number":25,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:36 {"time":"2020-03-29T14:47:34.3238673Z","log_number":26,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:36 {"time":"2020-03-29T14:47:34.3239869Z","log_number":27,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:36 {"time":"2020-03-29T14:47:34.324092Z","log_number":28,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:36 {"time":"2020-03-29T14:47:34.3241733Z","log_number":29,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:36 {"time":"2020-03-29T14:47:34.3242042Z","log_number":30,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:38 {"time":"2020-03-29T14:47:37.5479351Z","log_number":31,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:38 {"time":"2020-03-29T14:47:37.5479827Z","log_number":32,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:38 {"time":"2020-03-29T14:47:37.548019Z","log_number":33,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:38 {"time":"2020-03-29T14:47:37.5480518Z","log_number":34,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:38 {"time":"2020-03-29T14:47:37.5480878Z","log_number":35,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:43 {"time":"2020-03-29T14:47:40.807578Z","log_number":36,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:43 {"time":"2020-03-29T14:47:40.8076412Z","log_number":37,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:43 {"time":"2020-03-29T14:47:40.8078322Z","log_number":38,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:43 {"time":"2020-03-29T14:47:40.8078919Z","log_number":39,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:43 {"time":"2020-03-29T14:47:40.8079409Z","log_number":40,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:46 {"time":"2020-03-29T14:47:44.0584113Z","log_number":41,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:46 {"time":"2020-03-29T14:47:44.0585029Z","log_number":42,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:46 {"time":"2020-03-29T14:47:44.0586498Z","log_number":43,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:46 {"time":"2020-03-29T14:47:44.0588882Z","log_number":44,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:46 {"time":"2020-03-29T14:47:44.0589365Z","log_number":45,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:49 {"time":"2020-03-29T14:47:47.2840995Z","log_number":46,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:49 {"time":"2020-03-29T14:47:47.2841462Z","log_number":47,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:49 {"time":"2020-03-29T14:47:47.2841829Z","log_number":48,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:49 {"time":"2020-03-29T14:47:47.2842196Z","log_number":49,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:49 {"time":"2020-03-29T14:47:47.2842574Z","log_number":50,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:53 {"time":"2020-03-29T14:47:50.5083565Z","log_number":51,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:53 {"time":"2020-03-29T14:47:50.508478Z","log_number":52,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:53 {"time":"2020-03-29T14:47:50.5085626Z","log_number":53,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:53 {"time":"2020-03-29T14:47:50.508645Z","log_number":54,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:53 {"time":"2020-03-29T14:47:50.5087275Z","log_number":55,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:54 {"time":"2020-03-29T14:47:53.7376903Z","log_number":56,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:54 {"time":"2020-03-29T14:47:53.7378283Z","log_number":57,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:54 {"time":"2020-03-29T14:47:53.7378546Z","log_number":58,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:54 {"time":"2020-03-29T14:47:53.7378798Z","log_number":59,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:54 {"time":"2020-03-29T14:47:53.7379062Z","log_number":60,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:59 {"time":"2020-03-29T14:47:56.9641589Z","log_number":61,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:59 {"time":"2020-03-29T14:47:56.9642047Z","log_number":62,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:59 {"time":"2020-03-29T14:47:56.9642385Z","log_number":63,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:59 {"time":"2020-03-29T14:47:56.9642741Z","log_number":64,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:47:59 {"time":"2020-03-29T14:47:56.9643086Z","log_number":65,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:00.1846446Z","log_number":66,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:00.1849351Z","log_number":67,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:00.185024Z","log_number":68,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:00.1851058Z","log_number":69,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:00.185216Z","log_number":70,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:03.4149149Z","log_number":71,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:03.4149684Z","log_number":72,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:03.4150288Z","log_number":73,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:03.4150616Z","log_number":74,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:03 {"time":"2020-03-29T14:48:03.4150882Z","log_number":75,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:09 {"time":"2020-03-29T14:48:06.6305512Z","log_number":76,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:09 {"time":"2020-03-29T14:48:06.6306168Z","log_number":77,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:09 {"time":"2020-03-29T14:48:06.630654Z","log_number":78,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:09 {"time":"2020-03-29T14:48:06.6306901Z","log_number":79,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:09 {"time":"2020-03-29T14:48:06.630722Z","log_number":80,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:09.8811353Z","log_number":81,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:09.8811903Z","log_number":82,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:09.8813586Z","log_number":83,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:09.881461Z","log_number":84,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:09.8814917Z","log_number":85,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:13.108709Z","log_number":86,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:13.1087515Z","log_number":87,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:13.1087804Z","log_number":88,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:13.1088071Z","log_number":89,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:13 {"time":"2020-03-29T14:48:13.1088367Z","log_number":90,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:19 {"time":"2020-03-29T14:48:16.3398516Z","log_number":91,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:19 {"time":"2020-03-29T14:48:16.3399571Z","log_number":92,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:19 {"time":"2020-03-29T14:48:16.3400326Z","log_number":93,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:19 {"time":"2020-03-29T14:48:16.3401088Z","log_number":94,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:19 {"time":"2020-03-29T14:48:16.3401913Z","log_number":95,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:20 {"time":"2020-03-29T14:48:19.5732148Z","log_number":96,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:20 {"time":"2020-03-29T14:48:19.5732727Z","log_number":97,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:20 {"time":"2020-03-29T14:48:19.5733188Z","log_number":98,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:20 {"time":"2020-03-29T14:48:19.5734635Z","log_number":99,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:20 {"time":"2020-03-29T14:48:19.5735028Z","log_number":100,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:22.7980569Z","log_number":101,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:22.7981074Z","log_number":102,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:22.7981442Z","log_number":103,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:22.7981789Z","log_number":104,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:22.7982133Z","log_number":105,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:26.014402Z","log_number":106,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:26.0144424Z","log_number":107,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:26.0144691Z","log_number":108,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:26.0144952Z","log_number":109,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:26 {"time":"2020-03-29T14:48:26.0145219Z","log_number":110,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:32 {"time":"2020-03-29T14:48:29.2502572Z","log_number":111,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:32 {"time":"2020-03-29T14:48:29.2503874Z","log_number":112,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:32 {"time":"2020-03-29T14:48:29.2504211Z","log_number":113,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:32 {"time":"2020-03-29T14:48:29.2504582Z","log_number":114,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:32 {"time":"2020-03-29T14:48:29.2504939Z","log_number":115,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:35 {"time":"2020-03-29T14:48:32.4760146Z","log_number":116,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:35 {"time":"2020-03-29T14:48:32.4760785Z","log_number":117,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:35 {"time":"2020-03-29T14:48:32.4761333Z","log_number":118,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:35 {"time":"2020-03-29T14:48:32.4761878Z","log_number":119,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:35 {"time":"2020-03-29T14:48:32.4762447Z","log_number":120,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:36 {"time":"2020-03-29T14:48:35.6708842Z","log_number":121,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:36 {"time":"2020-03-29T14:48:35.670933Z","log_number":122,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:36 {"time":"2020-03-29T14:48:35.6709621Z","log_number":123,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:36 {"time":"2020-03-29T14:48:35.6709907Z","log_number":124,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:36 {"time":"2020-03-29T14:48:35.6710199Z","log_number":125,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:38.9061681Z","log_number":126,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:38.9062269Z","log_number":127,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:38.9062596Z","log_number":128,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:38.9062914Z","log_number":129,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:38.9063274Z","log_number":130,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:42.1344576Z","log_number":131,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:42.134504Z","log_number":132,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:42.1345343Z","log_number":133,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:42.1345651Z","log_number":134,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:43 {"time":"2020-03-29T14:48:42.1345989Z","log_number":135,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:46 {"time":"2020-03-29T14:48:45.3598228Z","log_number":136,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:46 {"time":"2020-03-29T14:48:45.3599542Z","log_number":137,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:46 {"time":"2020-03-29T14:48:45.3599961Z","log_number":138,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:46 {"time":"2020-03-29T14:48:45.3600404Z","log_number":139,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:46 {"time":"2020-03-29T14:48:45.3600832Z","log_number":140,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:49 {"time":"2020-03-29T14:48:48.5900569Z","log_number":141,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:49 {"time":"2020-03-29T14:48:48.5901107Z","log_number":142,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:49 {"time":"2020-03-29T14:48:48.5901475Z","log_number":143,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:49 {"time":"2020-03-29T14:48:48.5901815Z","log_number":144,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:48:49 {"time":"2020-03-29T14:48:48.5902176Z","log_number":145,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 {"time":"2020-03-29T14:49:56.2982689Z","log_number":246,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 err: index(146) != data.No(246), line:{"time":"2020-03-29T14:49:56.2982689Z","log_number":246,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 {"time":"2020-03-29T14:49:56.2983155Z","log_number":247,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 err: index(146) != data.No(247), line:{"time":"2020-03-29T14:49:56.2983155Z","log_number":247,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 {"time":"2020-03-29T14:49:56.2983447Z","log_number":248,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 err: index(146) != data.No(248), line:{"time":"2020-03-29T14:49:56.2983447Z","log_number":248,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 {"time":"2020-03-29T14:49:56.298379Z","log_number":249,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 err: index(146) != data.No(249), line:{"time":"2020-03-29T14:49:56.298379Z","log_number":249,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 {"time":"2020-03-29T14:49:56.2984133Z","log_number":250,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
2020/03/29 14:49:58 err: index(146) != data.No(250), line:{"time":"2020-03-29T14:49:56.2984133Z","log_number":250,"data":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
^C
```
