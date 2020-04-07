FROM golang:latest

ADD . /go/src/github.com/kunaljaydesai/webserver/iterative/
CMD go run /go/src/github.com/kunaljaydesai/webserver/iterative/iterative/main.go