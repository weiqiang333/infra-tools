#!/bin/bash

export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

go build infra-tools.go