#!/bin/sh
go get
env GOOS=windows GOARCH=amd64 go build -o bin/stroompunt-amd64.exe
env GOOS=linux GOARCH=amd64 go build -o bin/stroompunt-amd64.linux
env GOOS=darwin GOARCH=amd64 go build -o bin/stroompunt-amd64.darwin
