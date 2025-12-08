package models

// Query request to database
type Query struct {
	Command Command
	Args    []string
}

func NewQuery(command Command, args []string) *Query {
	return &Query{
		Command: command,
		Args:    args,
	}
}
