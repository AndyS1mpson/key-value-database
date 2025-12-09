package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"syscall"
	"time"

	"go.uber.org/zap"

	tcpClient "github.com/AndyS1mpson/key-value-database/internal/infrastructure/network/tcp_client"
	utils "github.com/AndyS1mpson/key-value-database/internal/utils/size_parser"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	logger, _ := zap.NewProduction()

	dbServer, err := initDBServerConnection()
	if err != nil {
		logger.Fatal("connection was closed", zap.Error(err))
	}

	if err = runCLIClient(dbServer, logger); err != nil {
		exitCode = failExitCode
	}

	return exitCode
}

func initDBServerConnection() (*tcpClient.TCPClient, error) {
	address := flag.String("address", "localhost:3223", "Address of the database")
	idleTimeout := flag.Duration("idle_timeout", time.Minute, "Idle timeout for connection")
	maxMessageSizeStr := flag.String("max_message_size", "4KB", "Max message size for connection")
	flag.Parse()

	bufferSize, err := utils.ParseSize(*maxMessageSizeStr)
	if err != nil {
		return nil, err
	}

	var options []tcpClient.Option

	options = append(options, tcpClient.WithClientIdleTimeout(*idleTimeout))
	options = append(options, tcpClient.WithClientBufferSize(uint(bufferSize)))

	client, err := tcpClient.NewTCPClient(*address, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to init tcp client: %w", err)
	}

	return client, nil
}

func runCLIClient(client *tcpClient.TCPClient, logger *zap.Logger) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("[key-value db] > ")
		request, err := reader.ReadString('\n')
		if errors.Is(err, syscall.EPIPE) {
			logger.Fatal("connection was closed", zap.Error(err))
		} else if err != nil {
			logger.Error("failed to read query", zap.Error(err))
		}

		response, err := client.Send([]byte(request))
		if errors.Is(err, syscall.EPIPE) {
			logger.Fatal("connection was closed", zap.Error(err))
		} else if err != nil {
			logger.Error("failed to send query", zap.Error(err))
		}

		fmt.Println(string(response))
	}
}
