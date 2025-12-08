package parser

import (
	"errors"
	"strings"

	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
	"go.uber.org/zap"
)

var (
	errEmptyQuery            = errors.New("empty query string")
	errUnknownCommand        = errors.New("unknown command")
	errIncorrectNumberOfArgs = errors.New("incorrect number of command arguments")
)

// Parse convert raw query to Query model
func (q *QueryParser) Parse(rawQuery string) (*models.Query, error) {
	tokens := strings.Fields(rawQuery)

	if len(tokens) == 0 {
		q.logger.Debug("empty tokens", zap.String("query", rawQuery))

		return nil, errEmptyQuery
	}

	command := models.Command(tokens[0])
	if !models.IsCommandExist(command) {
		q.logger.Debug("invalid command", zap.String("query", rawQuery))

		return nil, errUnknownCommand
	}

	query := models.NewQuery(command, tokens[1:])
	argsNumber := models.GetCommandArgsNumber(command)
	if len(query.Args) != argsNumber {
		q.logger.Debug("invalid arguments for query", zap.String("query", rawQuery))

		return nil, errIncorrectNumberOfArgs
	}

	return query, nil
}
