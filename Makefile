SERVER_APP_NAME=kv-db-server
CLI_APP_NAME=kb-db-cli

build-server:
	go build -o ${SERVER_APP_NAME} cmd/server/main.go

build-cli:
	go build -o ${CLI_APP_NAME} cmd/cli/main.go

run-server: build-server
	./${SERVER_APP_NAME}

run-cli: build-cli
	./${CLI_APP_NAME} $(ARGS)

run_unit_test:
	go test ./internal/...

run_test_coverage:
	go test ./... -coverprofile=coverage.out

lint:
	golangci-lint run
