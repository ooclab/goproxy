# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w
LDFLAGS=-ldflags "-s"
#LDFLAGS=

PROGRAM_NAME=goproxy

all:
	$(GOBUILD) -v $(LDFLAGS) -o $(PROGRAM_NAME)

install:
	$(GOINSTALL) -v
