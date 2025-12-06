
run-server:
	go run cmd/cli/main.go

# run linters
lint:
	golangci-lint run

test:
	go test -v ./...
