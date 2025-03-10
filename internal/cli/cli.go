package cli

import (
	"flag"
	"fmt"
)

// Options representa las opciones de la línea de comandos.
type Options struct {
	ShowVersion bool
	ShowHelp    bool
	NoPrefix    bool
	CustomPrefix string
}

// ParseFlags parsea los flags de la línea de comandos.
func ParseFlags() (Options, error) {
	opts := Options{}

	flag.BoolVar(&opts.ShowVersion, "version", false, "Show the version of the application")
	flag.BoolVar(&opts.ShowHelp, "help", false, "Show available flags")
	flag.BoolVar(&opts.NoPrefix, "no-prefix", false, "Disable the prefix in the commit message. [branch] by default")
	flag.StringVar(&opts.CustomPrefix, "prefix", "", "Define a custom prefix for the commit message")

	flag.Parse()

	return opts, nil
}

// PrintHelp muestra la ayuda de la aplicación.
func PrintHelp() {
	fmt.Println("Usage: gocommit [options]")
	flag.PrintDefaults()
}
