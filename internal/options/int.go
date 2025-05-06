package options

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/ayayaakasvin/random/internal/lib/logger"
	"github.com/ayayaakasvin/random/internal/lib/randomtool"
)

type IntCommand struct {
	Min, Max int64
}

func NewIntCommand () *IntCommand {
	return &IntCommand{}
}

func (i *IntCommand) Run () error {
	set := flag.NewFlagSet("integer", flag.ExitOnError)
	set.Usage = i.Help

	set.Int64Var(&i.Min, "min", 0, "Specify the minimum value for the range(default: 0)")
	set.Int64Var(&i.Max, "max", math.MaxInt64-1, "Specify the maximum value for the range(default: 9223372036854775806)")
	
	set.Parse(os.Args[2:])
	if i.Min >= i.Max {
        logger.ErrorLog("Invalid range: min must be less than max")
        return fmt.Errorf("invalid range: min (%d) must be less than max (%d)", i.Min, i.Max)
    }

	result, err := randomtool.RandomInt(i.Min, i.Max)
	if err != nil {
		logger.ErrorLog("%v", err)
		return err
	}

	fmt.Fprintf(os.Stdout, "%d\n", result)
	return nil
}

func (i *IntCommand) Help() {
	fmt.Fprintf(os.Stderr, `Usage of random integer [options]:
Generates random integer with specified options:

Options:
	-min	Specify the minimum value for the range (default: 0)
	-max	Specify the maximum value for the range (default: 9223372036854775806[math.MaxInt64-1])
`)
}