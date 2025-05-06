package options

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = Help
}

// RunApplication runs the app and parses flags
func RunApplication() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "No command specified")
		os.Exit(1)
	}

	mapOfCommands := GetSubcommandMap()

	command, ok := mapOfCommands[os.Args[1]]
	if !ok {
		flag.Usage()
		os.Exit(1)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

func Help() {
	fmt.Fprintf(os.Stderr, `Usage of random [options]
Options:
	string [options]	Generates random string with specified options
	integer [options]	Generates random integer with specified options
	float [options]		Generates random float with specified options
	Use random [options] --help to get information
`)
}