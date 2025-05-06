package options

type Subcommand interface {
	Run () error
	Help ()
}

// Returns Subcommands map
func GetSubcommandMap () map[string]Subcommand {
	var mapOfCommands map[string]Subcommand = make(map[string]Subcommand)

	mapOfCommands["string"] = NewStringCommand()
	mapOfCommands["integer"] = NewIntCommand()
	mapOfCommands["float"] = NewFloatCommand()

	return mapOfCommands
}
