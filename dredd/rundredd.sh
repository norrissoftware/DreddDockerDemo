#!/bin/bash

cd $GOPATH/src/github.com/jasonrichardsmith/DreddDockerDemo
glide install
go build main.go
dredd
