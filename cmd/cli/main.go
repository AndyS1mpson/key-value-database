package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AndyS1mpson/key-value-database/internal/container"
	"github.com/AndyS1mpson/key-value-database/internal/database"
	"go.uber.org/zap"
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

	serviceContainer, shutdown := container.NewContainer(container.LoadConfig())
	defer func() {
		if panicErr := recover(); panicErr != nil {
			exitCode = failExitCode
		}

		if err != nil {
			exitCode = failExitCode
		}
	}()
	defer shutdown()

	ctx, stop := signal.NotifyContext(serviceContainer.Ctx(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db := serviceContainer.GetDatabase()

	if err = runCLI(ctx, db, serviceContainer.GetLogger()); err != nil {
		exitCode = failExitCode
	}

	return exitCode
}

func runCLI(ctx context.Context, db *database.Database, logger *zap.Logger) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("[key-value db] > ")
		request, err := reader.ReadString('\n')
		if err != nil {
			logger.Error("failed to read query", zap.Error(err))

			return err
		}

		response := db.HandleQuery(ctx, request)

		fmt.Println(string(response))
	}
}
