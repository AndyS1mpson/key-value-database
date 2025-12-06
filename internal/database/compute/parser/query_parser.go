package parser

import "go.uber.org/zap"

// QueryParser parse database raw queries
type QueryParser struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *QueryParser {
	return &QueryParser{
		logger: logger,
	}
}
