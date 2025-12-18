package parser

import "go.uber.org/zap"

// QueryParser parses raw query strings into structured Query models.
type QueryParser struct {
	logger *zap.Logger
}

// New creates a new QueryParser instance with the given logger.
func New(logger *zap.Logger) *QueryParser {
	return &QueryParser{
		logger: logger,
	}
}
