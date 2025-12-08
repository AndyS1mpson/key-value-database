package container

import (
	"fmt"

	"go.uber.org/zap"
)

func (c *Container) GetLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("can not init logger: %s", err))
	}

	return logger
}
