package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lukecjohnson/dot/lib/commands"
)

func main() {
	// Arguments
	arguments := os.Args[1:]

	// Flags
	var v bool
	flag.BoolVar(&v,"v", false, "Prints the current installed version of this tool.")
	flag.BoolVar(&v,"version", false, "Prints the current installed version of this tool.")

	// Override default usage
	flag.Usage = commands.Help

	// Parse flags
	flag.Parse()

	// Too many arguments
	if len(arguments) > 1 {
		fmt.Println("Error: Too many arguments.")
		fmt.Println("Run \"dot -h\" for usage instructions.")
		os.Exit(2)
	}

	// No arguments - output usage instructions
	if len(arguments) == 0 {
		commands.Help()
	}

	// Run version command
	if v {
		commands.Version()
	}

	// Run root command
	script := arguments[0]
	commands.Root(script)
}
