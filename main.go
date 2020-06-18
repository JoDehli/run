package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

// Version : Dynamically set at build time to the most recent git tag
var Version = "DEV"

func main() {
  // Arguments
  arguments := os.Args[1:]

  // Flags
  versionFlag := flag.BoolP("version", "v", false, "Reports the current installed version of run.")
  listFlag := flag.BoolP("list", "l", false, "Lists all the available commands found in \"run.yaml\".")

  // Override default usage output
  flag.Usage = func() {
    fmt.Println("Usage:")
    fmt.Println("  run [command]")

    fmt.Println("\n" + "Other options:")
    flag.PrintDefaults()

    os.Exit(0)
  }

  // Parse flags
  flag.Parse()

  // Too many arguments
  if len(arguments) > 1 {
    fmt.Println("Error: too many arguments.")
    fmt.Println("Run \"run --help\" for usage instructions.")
    os.Exit(2)
  }

  // No arguments/flags - show usage instructions
  if len(arguments) == 0 {
    flag.Usage()
  }

  // Version flag - print version
  if *versionFlag {
    fmt.Println(Version)
    os.Exit(0)
  }

  // Check for "run.yaml"
  if _, err := os.Stat("run.yaml"); os.IsNotExist(err) {
    fmt.Println("Error: unable to resolve \"run.yaml\" in the current directory.")
    os.Exit(1)
  }

  // Read "run.yaml"
  data, err := ioutil.ReadFile("run.yaml")
  if err != nil {
    fmt.Println("Error: unable to read \"run.yaml\".")
    os.Exit(1)
  }

  // Parse "run.yaml"
  commands := make(map[string]string)
  if err := yaml.Unmarshal(data, &commands); err != nil {
    fmt.Println("Error: unable to parse \"run.yaml\".")
    os.Exit(1)
  }

  // List flag - print available commands
  if *listFlag {
    fmt.Println("Available commands:")
    for command := range commands {
      fmt.Println("  " + command)
    }

    os.Exit(0)
  }

  // Provided command
  command := arguments[0]

  // Check if command exists in "run.yaml"
  if _, ok := commands[command]; !ok {
    fmt.Printf("Error: command \"%s\" cannot be found in \"run.yaml\".\n", command)
    os.Exit(2)
  }

  // Execute command
  cmd := exec.Command("/bin/sh", "-c", commands[command])
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    fmt.Printf("Error: failed to execute \"%s\".\n", commands[command])
    os.Exit(1)
  }
}
