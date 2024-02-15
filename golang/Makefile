GOCMD=go
GOBUILD=$(GOCMD) build -ldflags="-s -w"
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_FOLDER=bin

build:
	$(GOCLEAN)
	$(GOBUILD) -a -o $(BINARY_FOLDER)/ .

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_FOLDER)/*

run:
	$(GORUN) . 

darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_FOLDER)/darwin/ip-watch-amd64 .
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BINARY_FOLDER)/darwin/ip-watch-arm64 .

windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_FOLDER)/windows/ .

linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_FOLDER)/linux/ .

build_all: clean darwin windows linux