# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODEV=bee run -main=cmd/users-service/main.go
GOLINT=golangci-lint run
GOGET=$(GOCMD) get
BINARY_NAME=users-service
BINARY_DIR=tmp
PACKAGE_NAME=greml1n/otus-microservices-hw02
DB_SERVICE=hw02_k8s_postgres
VERSION=v0.0.1

lint:
	$(GOLINT) ./...

clean:
	$(GOCLEAN)
	rm -r $(BINARY_DIR)

dev:
	$(GODEV)

db_setup:
	docker_compose -f deployments/docker-compose.yml up -d

db_up:
	docker run --name ${DB_SERVICE} -d -e POSTGRES_USER=superuser -e POSTGRES_PASSWORD=superpass postgres:13.3-alpine

db_down:
	docker stop ${DB_SERVICE} && docker rm ${DB_SERVICE}

run:
	./$(BINARY_DIR)/$(BINARY_NAME)

tidy:
	$(GOCMD) mod tidy

build:
	mkdir -p $(BINARY_DIR) && $(GOBUILD) -o $(BINARY_DIR) ./...

push:
	docker build -t $(PACKAGE_NAME):$(VERSION) -t $(PACKAGE_NAME):latest .
	docker push -a $(PACKAGE_NAME)