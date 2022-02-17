GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=hello-go

build:
	(cd cmd/hello; $(GOBUILD) -o ../../bin/$(BINARY_NAME))

run:
	./bin/$(BINARY_NAME)

test:
	(cd cmd/hello; $(GOTEST))

