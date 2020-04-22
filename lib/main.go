package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Arguments
	arguments := os.Args[1:]

	// Flags
	v := flag.Bool("v", false, "Prints the current installed version of this tool.")

	// Parse flags
	flag.Parse()

	// Too many arguments
	if len(arguments) > 1 {
		fmt.Println("Too many arguments.")
		fmt.Println("Run \"do -h\" for usage instructions.")
		os.Exit(2)
	}

	// No arguments - output usage instructions
	if len(arguments) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	// Run version command
	if *v {
		fmt.Println("version")
		os.Exit(0)
	}

	script := arguments[0]
	fmt.Println(script)
}
