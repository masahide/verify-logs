# verify-logs



[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/masahide/verify-logs)](https://hub.docker.com/repository/docker/masahide/verify-logs)[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/masahide/verify-logs)](https://hub.docker.com/r/masahide/verify-logs/builds)


## usage

```bash
# use ~/.aws/config 
docker run -e AWS_PROFILE=profile-name -e AWS_SDK_LOAD_CONFIG=1 -e AWS_DEFAULT_REGION=ap-northeast-1 -v $HOME:/root masahide/verify-logs

# set loggroup-name
docker run -e LogGroupName=loggroup-name -e AWS_PROFILE=profile-name -e AWS_SDK_LOAD_CONFIG=1 -e AWS_DEFAULT_REGION=ap-northeast-1 -v $HOME:/root masahide/verify-logs
```
