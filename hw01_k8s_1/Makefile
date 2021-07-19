# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODEV=air
GOLINT=golangci-lint run
GOGET=$(GOCMD) get
BINARY_NAME=hello-service
BINARY_DIR=tmp

lint:
	$(GOLINT) ./...

clean:
	$(GOCLEAN)
	rm -r $(BINARY_DIR)

dev:
	$(GODEV)

run:
	./$(BINARY_DIR)/$(BINARY_NAME)

tidy:
	$(GOCMD) mod tidy

build:
	mkdir -p $(BINARY_DIR) && $(GOBUILD) -o $(BINARY_DIR) ./...