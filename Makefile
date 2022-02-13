GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=hello-go

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) cmd/hello/main.go

run:
	./bin/$(BINARY_NAME)
