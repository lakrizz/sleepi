#!/bin/bash

rm main 2> /dev/null
go mod vendor
go build cmd/sleepi/main.go

./main

