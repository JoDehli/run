package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

// Dynamically set at build time to the most recent git tag
var currentVersion = "DEV"

func main() {
	// Command line flags
	version := flag.BoolP("version", "v", false, "Display the current version of run")
	list := flag.BoolP("list", "l", false, "List all the available commands found in \"run.yaml\"")

	// Override default usage function
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  run [command]")

		fmt.Println("\n" + "Other options:")
		flag.PrintDefaults()

		os.Exit(0)
	}

	// Parse flags
	flag.Parse()

	// Parse non-flag arguments
	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
	}

	if len(args) > 1 {
		fmt.Println("Error: too many arguments.")
		fmt.Println("Run \"run --help\" for usage instructions.")
		os.Exit(2)
	}

	if *version {
		fmt.Println(currentVersion)
		os.Exit(0)
	}

	commandsFile := "run.yaml"
	if _, err := os.Stat(commandsFile); os.IsNotExist(err) {
		commandsFile = "run.yml"
		if _, err := os.Stat(commandsFile); os.IsNotExist(err) {
			fmt.Println("Error: unable to find \"run.yaml\" in the current directory.")
			os.Exit(1)
		}
	}

	data, err := ioutil.ReadFile(commandsFile)
	if err != nil {
		fmt.Println("Error: unable to read \"run.yaml\".")
		os.Exit(1)
	}

	commands := make(map[string]string)
	if err := yaml.Unmarshal(data, &commands); err != nil {
		fmt.Println("Error: unable to parse \"run.yaml\".")
		os.Exit(1)
	}

	if *list {
		fmt.Println("Available commands:")
		for command := range commands {
			fmt.Println("  " + command)
		}

		os.Exit(0)
	}

	command := args[0]

	if _, ok := commands[command]; !ok {
		fmt.Printf("Error: command \"%s\" cannot be found in \"run.yaml\".\n", command)
		os.Exit(2)
	}

	cmd := exec.Command("/bin/sh", "-c", commands[command])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: failed to execute \"%s\".\n", commands[command])
		os.Exit(1)
	}
}
