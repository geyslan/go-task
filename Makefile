.PHONY: build-containers
build-containers:
	@docker-compose build

.PHONY: start-containers
start-containers:
	@docker-compose up

.PHONY: start-containers-daemon
start-containers-daemon:
	@docker-compose up -d

.PHONY: stop-containers
stop-containers:
	@docker-compose down

.PHONY: build-proto
build-proto:
	@protoc --go-grpc_out=serverDomain/protos --go_out=serverDomain/protos serverDomain/protos/port.proto

.PHONY: run-tests
run-tests:
	@go test -cover ./...
