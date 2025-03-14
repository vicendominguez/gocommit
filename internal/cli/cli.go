package cli

import (
	"flag"
	"fmt"
)

// Options represents the command line options.
type Options struct {
	ShowVersion  bool
	ShowHelp     bool
	NoPrefix     bool
	CustomPrefix string
	DryRun       bool
}

func ParseFlags() (Options, error) {
	opts := Options{}

	flag.BoolVar(&opts.ShowVersion, "version", false, "Show the version of the application")
	flag.BoolVar(&opts.ShowHelp, "help", false, "Show available flags")
	flag.BoolVar(&opts.DryRun, "dry-run", false, "Show a message example")
	flag.BoolVar(&opts.NoPrefix, "no-prefix", false, "Disable the prefix in the commit message. [branch] by default")
	flag.StringVar(&opts.CustomPrefix, "prefix", "", "Define a custom prefix for the commit message")

	flag.Parse()

	return opts, nil
}

// PrintHelp show the help
func PrintHelp() {
	fmt.Println("Usage: gocommit [options]")
	flag.PrintDefaults()
}
