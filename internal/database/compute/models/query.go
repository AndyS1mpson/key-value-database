package models

// Query represents a parsed database request with command and arguments.
type Query struct {
	Command Command
	Args    []string
}

// NewQuery creates a new Query instance with the given command and arguments.
func NewQuery(command Command, args []string) *Query {
	return &Query{
		Command: command,
		Args:    args,
	}
}
