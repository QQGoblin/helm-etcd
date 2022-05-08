#!/bin/bash

export HTTPS_PROXY=http://172.28.235.169:10889
export HTTP_PROXY=http://172.28.235.169:10889

GOOS=linux GOARCH=amd64 CGO_ENABLED=0  GO111MODULE=on GOPROXY=https://goproxy.cn go mod tidy
GOOS=linux GOARCH=amd64 CGO_ENABLED=0  GO111MODULE=on GOPROXY=https://goproxy.cn go mod vendor
GOOS=linux GOARCH=amd64 CGO_ENABLED=0  GO111MODULE=on GOPROXY=https://goproxy.cn go build -o ./bin/helm-etcd -ldflags "" ./main.go
