FROM golang:1.17.6-alpine3.15

USER root

RUN mkdir -p /app /toolbox

RUN apk add --no-cache openssh shadow git grep

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/todo-checker .

WORKDIR /app

