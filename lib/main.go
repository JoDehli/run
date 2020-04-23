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

	// Version flag
	var v bool
	flag.BoolVar(&v,"v", false, "Prints the current installed version of this tool.")
	flag.BoolVar(&v,"version", false, "Prints the current installed version of this tool.")

	// Scripts flag
	var s bool
	flag.BoolVar(&s,"s", false, "Prints all the available scripts found in \"scripts.yaml\".")
	flag.BoolVar(&s,"scripts", false, "Prints all the available scripts found in \"scripts.yaml\".")

	// Override default usage function
	flag.Usage = commands.Help

	// Parse flags
	flag.Parse()

	// Too many arguments
	if len(arguments) > 1 {
		fmt.Println("Error: Too many arguments.")
		fmt.Println("Run \"dot -h\" for usage instructions.")
		os.Exit(2)
	}

	// No arguments - run help command
	if len(arguments) == 0 {
		commands.Help()
	}

	// Run version command
	if v {
		commands.Version()
	}

	// Run scripts command
	if s {
		commands.Scripts()
	}

	// Run root command
	script := arguments[0]
	commands.Root(script)
}
