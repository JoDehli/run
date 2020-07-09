package main

import (
	"errors"
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

	if *version {
		fmt.Println(currentVersion)
		os.Exit(0)
	}

	if *list {
		if err := printAvailableCommands(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Parse commands (non-flag arguments)
	commands := flag.Args()

	if len(commands) == 0 {
		if err := printAvailableCommands(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	availableCommands, err := getAvailableCommands()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Ensure all provided commands exist before trying to execute
	for _, c := range commands {
		if _, ok := availableCommands[c]; !ok {
			fmt.Printf("command \"%s\" could not be found\n", c)
			os.Exit(1)
		}
	}

	// Execute each provided command
	for _, c := range commands {
		if err := executeCommand(availableCommands[c]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func getAvailableCommands() (map[string]string, error) {
	file := "run.yaml"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		file = "run.yml"
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return nil, errors.New("unable to find \"run.yaml\" in the current directory")
		}
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.New("unable to read \"run.yaml\"")
	}

	commands := make(map[string]string)
	if err := yaml.Unmarshal(data, &commands); err != nil {
		return nil, errors.New("unable to parse \"run.yaml\"")
	}

	return commands, nil
}

func printAvailableCommands() error {
	availableCommands, err := getAvailableCommands()
	if err != nil {
		return err
	}

	fmt.Println("\nAvailable commands:")
	for c := range availableCommands {
		fmt.Printf("- %s\n", c)
	}
	fmt.Println()

	return nil
}

func executeCommand(command string) error {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println()
		return errors.New("failed to execute command")
	}

	return nil
}
