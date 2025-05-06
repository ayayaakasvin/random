package options

import (
	"flag"
	"fmt"
	"os"

	"github.com/ayayaakasvin/random/internal/lib/logger"
	"github.com/ayayaakasvin/random/internal/lib/randomtool"
)

type FloatCommand struct {
	Min, Max float64
	Precision uint
}

func NewFloatCommand () *FloatCommand {
	return &FloatCommand{}
}

func (f *FloatCommand) Run () error {
	set := flag.NewFlagSet("float", flag.ExitOnError)
	set.Usage = f.Help

	set.Float64Var(&f.Min, "min", 0.0, "Specify the minimum value for the range(default: 0)")
	set.Float64Var(&f.Max, "max", 1e6, "Specify the maximum value for the range(default: 1000000)")
	set.UintVar(&f.Precision, "precision", 6, "Specify the value for the precision(default: 6)")
	set.UintVar(&f.Precision, "p", 6, "Alias for -precision")

	set.Parse(os.Args[2:])
	if f.Min >= f.Max {
        logger.ErrorLog("Invalid range: min must be less than max")
        return fmt.Errorf("invalid range: min (%f) must be less than max (%f)", f.Min, f.Max)
    }

	result, err := randomtool.RandomFloat(f.Min, f.Max, f.Precision)
	if err != nil {
		logger.ErrorLog("Error during process: %v", err)
		return err
	}

	fmt.Fprintf(os.Stdout, "%f\n", result)
	return nil
}

func (f *FloatCommand) Help() {
	fmt.Fprintf(os.Stderr, `Usage of random float [options]:
Generates random float with specified options:

Options:
  	-min	Specify the minimum value for the range (default: 0)
  	-max	Specify the maximum value for the range (default: 9223372036854775806[math.MaxInt64-1])
	-precision	Specify the number of decimal places (default: 6)
  	-p	Alias for -precision
`)
}