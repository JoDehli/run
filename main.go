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

	// Print version and exit if --version flag was provided
	if *version {
		fmt.Println(currentVersion)
		os.Exit(0)
	}

	// Check available commands if --list flag was provided
	if *list {
		if err := printAvailableCommands(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Parse commands (non-flag arguments)
	commands := flag.Args()

	// Print available commands if no arguments were provided
	if len(commands) == 0 {
		if err := printAvailableCommands(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Get available commands found in run.yaml
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

	fmt.Println()
}

// Returns a parsed map of the commands found in run.yaml
func getAvailableCommands() (map[string]string, error) {
	// Check if run.yaml exists
	file := "run.yaml"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		// If run.yaml does not exit, check for run.yml
		file = "run.yml"
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return nil, errors.New("unable to find \"run.yaml\" in the current directory")
		}
	}

	// Read run.yaml
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.New("unable to read \"run.yaml\"")
	}

	// Parse run.yaml
	commands := make(map[string]string)
	if err := yaml.Unmarshal(data, &commands); err != nil {
		return nil, errors.New("unable to parse \"run.yaml\"")
	}

	// Returned parsed commands
	return commands, nil
}

// Prints a list of available commands
func printAvailableCommands() error {
	// Get available commands found in run.yaml
	availableCommands, err := getAvailableCommands()
	if err != nil {
		return err
	}

	// Print name of each command
	fmt.Println("\nAvailable commands:")
	for c := range availableCommands {
		fmt.Printf("- %s\n", c)
	}
	fmt.Println()

	// No error
	return nil
}

// Executes provided command
func executeCommand(command string) error {
	// Print shell command being executed
	fmt.Printf("\n\033[2m%s\033[0m\n", command)

	// Execute command
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println()
		return errors.New("failed to execute command")
	}

	// No error
	return nil
}
