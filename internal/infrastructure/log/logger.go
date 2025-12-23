package log

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	debugLevel = "debug"
	infoLevel  = "info"
	warnLevel  = "warn"
	errorLevel = "error"
)

const (
	defaultEncoding   = "json"
	defaultLevel      = zapcore.InfoLevel
	defaultOutputPath = "kv_db.log"
)

func ConfigureZapLogger(config *Config) (*zap.Logger, error) {
	level := defaultLevel
	output := defaultOutputPath

	if config != nil {
		if config.Level != "" {
			supportedLoggingLevels := map[string]zapcore.Level{
				debugLevel: zapcore.DebugLevel,
				infoLevel:  zapcore.InfoLevel,
				warnLevel:  zapcore.WarnLevel,
				errorLevel: zapcore.ErrorLevel,
			}

			var found bool
			if level, found = supportedLoggingLevels[config.Level]; !found {
				return nil, errors.New("logging level is incorrect")
			}
		}

		output = config.Output
	}

	loggerCfg := zap.Config{
		Encoding:    defaultEncoding,
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{output},
	}

	return loggerCfg.Build()
}
