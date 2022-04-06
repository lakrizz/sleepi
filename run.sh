#!/bin/bash

go mod tidy
go run cmd/$1/main.go
