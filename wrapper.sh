#!/bin/sh
go run  /go/src/github.com/golang/parse_article/grpc/server/server.go & 
go run  /go/src/github.com/golang/parse_article/grpc/http_client.go