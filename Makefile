# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w
LDFLAGS=-ldflags "-s"
STATIC_LDFLAGS=-a -installsuffix cgo -ldflags "-s -X main.buildstamp=`date -u '+%Y-%m-%dT%H:%M:%SZ'` -X main.githash=`git rev-parse HEAD | cut -c1-8`"

PROGRAM_NAME=goproxy

all:
	$(GOBUILD) -v $(LDFLAGS) -o $(PROGRAM_NAME)

install:
	$(GOINSTALL) -v

build-static:
	CGO_ENABLED=0 $(GOBUILD) -v $(STATIC_LDFLAGS) -o $(PROGRAM_NAME)
