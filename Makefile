GITTAG=`git rev-parse --short HEAD`
BUILD_TIME=`date -u +%Y.%m.%d-%H:%M:%S%Z`
VERSION=0.0.1
GOPATH ?= $(shell go env GOPATH)
GOFLAGS=-ldflags "-X "github.com/onerepo/faga/internal/version".GitCommit=${GITTAG} -X "github.com/onerepo/faga/internal/version".BuildTime=${BUILD_TIME} -X "github.com/onerepo/faga/internal/version".Version=${VERSION}"

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
CURDIR := $(shell pwd)
