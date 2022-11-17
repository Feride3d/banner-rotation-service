BIN := "./bin/banner-rotation-service"
DOCKER_IMG="banner-rotation-service:main"
MIGRATION := "/migration"
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd

build-img:
	docker build \
		-t $(DOCKER_IMG) \

run-img: build-img
	docker run $(DOCKER_IMG)

.PHONY: rabbit

rabbit:
    ## http://localhost:15672/ guest:guest
	docker run -d --name banner-stats -p 5672:5672 rabbitmq:3-management

test:
	go test -race -count 100 ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

proto:
	protoc --go_out=$(GOPATH)/src \
		--go_opt=paths=import \
		--go-grpc_out=$(GOPATH)/src \
		--go-grpc_opt=paths=import \
		third_party/google/api/annotations.proto \
		third_party/google/api/http.proto \
		third_party/google/protobuf/descriptor.proto \
		api/rotator/rotator.proto

generate: mkdir -p internal/pb
	protoc -I . \
    --go_out ./internal/pb/ --go_opt paths=source_relative \
    --go-grpc_out ./internal/pb/ --go-grpc_opt paths=source_relative \
    api/rotator/rotator.proto

generate-gateway: generate
	protoc -I . --grpc-gateway_out ./internal/pb \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		api/rotator/rotator.proto

dev-build-container:
	docker rm --force rotator-br
	docker rm --force postgres-br
	docker rm --force rabbit-br
	docker-compose build --no-cache

run:
	docker-compose up -d

down:
	docker-compose down

dev: dev-build-container
	docker-compose up

up:
	docker-compose up --build

migration-up:
	goose -dir {{.MIGRATION}} up

migration-down:
	goose -dir {{.MIGRATION}} down

.PHONY: build build-img run-img test lint generate-gateway run stop dev up integration-tests