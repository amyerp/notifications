# notifications Microservice

Full API Documentation in docs/ folder

## Build Microservoce

```
docker build --no-cache -t notifications:latest -f Dockerfile .
```
or
```
docker build -t notifications:latest -f Dockerfile .
```


## Run Microservice in Docker (in case if it in the same area with API Gateway)

```
docker run --name notifications \
--restart=always \
-v $(pwd)/config:/var/gufo/config \
-v $(pwd)/lang:/var/gufo/lang \
-v $(pwd)/templates:/var/gufo/templates \
-v $(pwd)/logs:/var/gufo/log \
-v $(pwd)/files:/var/gufo/files \
--network="gufo" \
-d notifications:latest
```

Before run microservice need to add in API Gateway config next lines

```
[microservices]
[microservices.notifications]
type = 'server'
host = 'notifications'
port = '5300'
entrypointversion = '1.0.0'
cron = 'false'

[microservices.notifications.matrix]
host = 'https://im.mymates.gmbh'
uname = 'amy'
pass = ''
token = ''
```
