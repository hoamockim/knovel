GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
APP=tasks
update:
	go mod tidy
all:
	make build 
build:
	$(GOBUILD) -v -ldflags="-extldflags=-static" -o $(APP) .
test:
	$(GOTEST) -v ./...  -covermode=count -coverprofile=sample.cov
run:
	./tasks
