package options

import (
	"flag"
	"fmt"
	"os"

	"github.com/ayayaakasvin/random/internal/lib/logger"
	"github.com/ayayaakasvin/random/internal/lib/randomtool"
)

type StringCommand struct {
	Length 			int
	AllowedSets 	[]string

}

func NewStringCommand () *StringCommand {
	return &StringCommand{}
}

func (s *StringCommand) Run() error {
	set := flag.NewFlagSet("string", flag.ExitOnError)
	set.Usage = s.Help

	set.IntVar(&s.Length, "length", 8, "Length of random string(default: 8)")
	set.IntVar(&s.Length, "len", 8, "Alias for length flag")

	var UpperNotAllowed bool
	set.BoolVar(&UpperNotAllowed, "no-upper", false, "Excludes uppercase letters from random string")
	set.BoolVar(&UpperNotAllowed, "u", false, "Alias for no-upper flag")

	var LowerNotAllowed bool
	set.BoolVar(&LowerNotAllowed, "no-lower", false, "Excludes lowercase letters from random string")
	set.BoolVar(&LowerNotAllowed, "l", false, "Alias for no-lower flag")

	var DigitsAllowed bool
	set.BoolVar(&DigitsAllowed, "digits", false, "Includes digits into random string")
	set.BoolVar(&DigitsAllowed, "d", false, "Alias for digits flag")

	var SpecialAllowed bool
	set.BoolVar(&SpecialAllowed, "special", false, "Includes special symbols into random string")
	set.BoolVar(&SpecialAllowed, "s", false, "Alias for special flag")

	set.Parse(os.Args[2:])

	if s.Length < 1 {
		logger.ErrorLog("Length can not be 0 or negative!")
		return fmt.Errorf("invalid flag value: %s=%d", "-length", s.Length) 
	}
	if !UpperNotAllowed {
		s.AllowedSets = append(s.AllowedSets, "upper")
	}
	if !LowerNotAllowed {
		s.AllowedSets = append(s.AllowedSets, "lower")
	}
	if DigitsAllowed {
		s.AllowedSets = append(s.AllowedSets, "digits")
	}
	if SpecialAllowed {
		s.AllowedSets = append(s.AllowedSets, "special")
	}
	if len(s.AllowedSets) == 0 {
		logger.ErrorLog("No set included. Exiting application...")
	}

	result, err := randomtool.RandomString(s.Length, s.AllowedSets)
	if err != nil {
		logger.ErrorLog("Error during process")
		logger.ErrorLog("Random string(not complete):")
		fmt.Fprintf(os.Stdout, "%s", result)
		return err
	}

	fmt.Fprintf(os.Stdout, "%s\n", result)

	return nil
}

func (s *StringCommand) Help() {
	fmt.Fprintf(os.Stderr, `Usage: random string [options]

Generates a random string with the specified options.

Options:
	-length, -l         Length of the string (default: 8)
	-no-upper, -u       Exclude uppercase letters
	-no-lower, -l       Exclude lowercase letters
	-digits, -d         Include digits
	-special, -s        Include special characters
`)
}