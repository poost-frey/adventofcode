# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

GOOS=windows
GOARCH=amd64

all: run-01b
build-01a: 
		GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o ./cmd/01a/bin/01a -v ./cmd/01a/*.go 
run-01a: build-01a
		./cmd/01a/bin/01a
build-01b: 
		GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o ./cmd/01b/bin/01b -v ./cmd/01b/*.go 
run-01b: build-01b
		./cmd/01b/bin/01b
build-02a: 
		GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o ./cmd/02a/bin/02a -v ./cmd/02a/*.go 
run-02a: build-02a
		./cmd/02a/bin/02a
clean: 
		$(GOCLEAN)
		rm -f ./cmd/*/bin/*
deps:
		dep ensure -v
fmt:
		$(GOFMT) github.com/poost-frey/adventofcode/...
test: 
		$(GOTEST) -v ./...
