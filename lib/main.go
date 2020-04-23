package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lukecjohnson/dot/lib/commands"
)

func main() {
	arguments := os.Args[1:]

	var versionFlag bool
	flag.BoolVar(&versionFlag, "v", false, "Prints the current installed version of this tool.")
	flag.BoolVar(&versionFlag, "version", false, "Prints the current installed version of this tool.")

	var listFlag bool
	flag.BoolVar(&listFlag, "l", false, "Prints all the available commands found in \"dot.yaml\".")
	flag.BoolVar(&listFlag, "list", false, "Prints all the available commands found in \"dot.yaml\".")

	flag.Usage = commands.Help
	
	flag.Parse()

	if len(arguments) > 1 {
		fmt.Println("Error: Too many arguments.")
		fmt.Println("Run \"dot -h\" for usage instructions.")
		os.Exit(2)
	}

	if len(arguments) == 0 {
		commands.Help()
	}

	if versionFlag {
		commands.Version()
	}

	if listFlag {
		commands.List()
	}

	command := arguments[0]
	commands.Run(command)
}
