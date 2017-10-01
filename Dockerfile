FROM golang:1.8-alpine
COPY . /usr/src/md-exporter
RUN  go build -o /usr/local/bin/md-exporter /usr/src/md-exporter/md-exporter.go
CMD  ["/usr/local/bin/md-exporter"]
