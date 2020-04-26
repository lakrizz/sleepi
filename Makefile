
PROJECTNAME=$(shell basename "$(PWD)")
# Go related variables.
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

MAKEFLAGS += --silent

exec:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(run)
