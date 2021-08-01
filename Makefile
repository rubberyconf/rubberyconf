
# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool

# Binary names
BINARY_NAME=rubberyconf
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

test:
	$(GOTEST) -v -coverpkg=./... -coverprofile=profile.cov ./...

coverage:
	$(GOTOOL) cover -func profile.cov

testcoverage: test coverage

build:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v cmd/server/main.go

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/github.com/rubberyconf/rubberyconf golang:latest go build -o "$(BINARY_UNIX)" -v