package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lukecjohnson/dot/lib/commands"
)

func main() {
	arguments := os.Args[1:]

	var v bool
	flag.BoolVar(&v, "v", false, "Prints the current installed version of this tool.")
	flag.BoolVar(&v, "version", false, "Prints the current installed version of this tool.")

	var l bool
	flag.BoolVar(&l, "l", false, "Prints all the available scripts found in \"scripts.yaml\".")
	flag.BoolVar(&l, "list", false, "Prints all the available scripts found in \"scripts.yaml\".")

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

	if v {
		commands.Version()
	}

	if l {
		commands.List()
	}

	script := arguments[0]
	commands.Run(script)
}
