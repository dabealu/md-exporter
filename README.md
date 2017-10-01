#### Simple software RAID (md) state exporter for prometheus
In almost any case, it's better to use **node-exporter** which can server md metrics and many more.  
Metric values: 1 - clean, 0 - degraded.  
Some parameters can be defined through env variables:  
- MD_URL - URL under which metrics will be available (default: /metrics)  
- MD_LISTEN - port which exporter listen (default: 8080)  
- MD_SYSPATH - path to /sys filesystem (default: /sys)  
- MD_HOST - values of hostname label (default: host's hostname)  

#### Build
binary:  
```
docker run --rm \
  -v "$PWD":/usr/src/md-exporter \
  -w /usr/src/md-exporter \
  golang:1.8-stretch go build -v
```
container with binary:  
```
docker build -t md-exporter .
```
(build image from scratch or alpine adding just compiled binary, to get minimal sized image)  

#### Run
From previously created image:
```
docker run -d \
  --name md-exporter \
  --hostname $HOSTNAME \
  -v /sys:/host-sys \
  -e MD_SYSPATH=/host-sys \
  -p 8080:8080 md-exporter
```
