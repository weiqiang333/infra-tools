#!/bin/bash

export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

go build -ldflags "-s -w"  -o ./bin/infra-tools infra-tools.go
go build -ldflags "-s -w"  -o ./bin/check cmd/check/check.go