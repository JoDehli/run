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
	list := flag.BoolP("list", "l", false, "Print all the available commands found in \"run.yaml\"")
	version := flag.BoolP("version", "v", false, "Print the current version of run")

	// Override default usage function
	flag.Usage = func() {
		fmt.Printf("\n%s\n\n", "Usage:")
		fmt.Println("  run [command] ...")

		fmt.Printf("\n\n%s\n\n", "Flags:")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("  -%s, --%s\t\t%s\n\n", f.Shorthand, f.Name, f.Usage)
		})

		os.Exit(0)
	}

	// Parse flags
	flag.Parse()

	// Print version and exit if --version flag was provided
	if *version {
		fmt.Println(currentVersion)
		os.Exit(0)
	}

	// Print available commands if --list flag was provided
	if *list {
		if err := printAvailableCommands(); err != nil {
			printError(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Parse commands (non-flag arguments)
	commands := flag.Args()

	// If no commands were provided, set default command
	if len(commands) == 0 {
		commands = []string{"default"}
	}

	// Execute commands
	if err := executeCommands(commands); err != nil {
		printError(err.Error())
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

	// No errors
	return nil
}

// Executes each provided command
func executeCommands(commands []string) error {

	// Get available commands found in run.yaml
	availableCommands, err := getAvailableCommands()
	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	// Ensure all provided commands exist before trying to execute
	for _, c := range commands {
		if _, ok := availableCommands[c]; !ok {
			printError(fmt.Sprintf("command \"%s\" could not be found", c))
			os.Exit(1)
		}
	}

	// Execute each provided command
	for _, c := range commands {

		// Print shell command being executed
		fmt.Printf("\n\033[2m%s\033[0m\n", availableCommands[c])

		// Execute command
		cmd := exec.Command("/bin/sh", "-c", availableCommands[c])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println()
			return fmt.Errorf("failed to execute command \"%s\"", c)
		}
	}

	// No errors
	return nil
}

// Prints formatted error message ("• Error: {message}")
func printError(msg string) {
	fmt.Printf("\n\033[1;31m•\033[0m Error: %s\n\n", msg)
}
