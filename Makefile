# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

GOOS=windows
GOARCH=amd64

all: run
build: 
		GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o ./cmd/04a/bin/04a -v ./cmd/04a/*.go 
run: build
		./cmd/04a/bin/04a
clean: 
		$(GOCLEAN)
		rm -f ./cmd/*/bin/*
deps:
		dep ensure -v
fmt:
		$(GOFMT) github.com/poost-frey/adventofcode/...
test: 
		$(GOTEST) -v ./...
