FROM golang:latest

MAINTAINER Giannicola Olivadoti <olinicola@gmail.com>

ADD . /go/src/github.com/oli-g/devops-welcome-task-app

RUN go install github.com/oli-g/devops-welcome-task-app/welcome

ENTRYPOINT ["welcome"]

EXPOSE 8080