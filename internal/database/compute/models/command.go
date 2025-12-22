package models

var (
	CommandUnknown Command = "UNKNOWN" // Unknown database command
	CommandGet     Command = "GET"     // Get value by key
	CommandSet     Command = "SET"     // Set key-value pair
	CommandDel     Command = "DEL"     // Delete key-value pair
)

var availableCommands = map[Command]struct{}{
	CommandGet: {},
	CommandSet: {},
	CommandDel: {},
}

// Command represents a supported database command type.
type Command string

// IsCommandExist checks if the given command is a valid supported command.
func IsCommandExist(command Command) bool {
	_, ok := availableCommands[command]

	return ok
}

var (
	getCommandArgsNumber int = 1
	setCommandArgsNumber int = 2
	delCommandArgsNumber int = 1
)

var commandArgsNumber = map[Command]int{
	CommandGet: getCommandArgsNumber,
	CommandSet: setCommandArgsNumber,
	CommandDel: delCommandArgsNumber,
}

// GetCommandArgsNumber returns the required number of arguments for a given command.
func GetCommandArgsNumber(command Command) int {
	return commandArgsNumber[command]
}
