package main

import (
	"context"
	"os"

	"github.com/AndyS1mpson/key-value-database/internal/container"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	var err error

	config, err := container.NewConfig()
	if err != nil {
		return failExitCode
	}

	serviceContainer, shutdown := container.NewContainer(*config)
	defer func() {
		if panicErr := recover(); panicErr != nil {
			exitCode = failExitCode
		}

		if err != nil {
			exitCode = failExitCode
		}
	}()
	defer shutdown()

	server := serviceContainer.GetTCPServer()
	database := serviceContainer.GetDatabase()

	server.HandleQueries(serviceContainer.Ctx(), func(ctx context.Context, query []byte) []byte {
		response := database.HandleQuery(ctx, string(query))
		return []byte(response)
	})

	return successExitCode
}
