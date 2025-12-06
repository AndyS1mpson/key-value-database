package models

var (
	CommandUnknown Command = "UNKNOWN" // unknown kv db command
	CommandGet     Command = "GET"     // get data from kv db
	CommandSet     Command = "SET"     // set data in kv db
	CommandDel     Command = "DEL"     // del data from kv db
)

var availableCommands = map[Command]struct{}{
	CommandGet: {},
	CommandSet: {},
	CommandDel: {},
}

// Command supported database command
type Command string

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

// GetCommandArgsNumber get number of command arguments
func GetCommandArgsNumber(command Command) int {
	return commandArgsNumber[command]
}
