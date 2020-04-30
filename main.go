package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
)

// Dynamically set at build time to the most recent git tag
var Version = "DEV"

func main() {
  // Arguments
  arguments := os.Args[1:]

  // Flags
  versionFlag := flag.Bool("version", false, "Reports the current installed version of run.")
  listFlag := flag.Bool("list", false, "Lists all the available commands found in \"run.json\".")

  // Override default usage output
  flag.Usage = func() {
    fmt.Println("Usage:")
    fmt.Println("  run <command>" + "\t\t" + "Runs the provided command found in \"run.json\".")

    fmt.Println("\n" + "Other options:")
    flag.VisitAll(func(f *flag.Flag) {
      fmt.Println("  --" + f.Name + "\t\t" + f.Usage)
    })

    os.Exit(0)
  }

  // Parse flags
  flag.Parse()

  // Too many arguments
  if len(arguments) > 1 {
    fmt.Println("Error: Too many arguments.")
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

  // Check for "run.json"
  if _, err := os.Stat("run.json"); os.IsNotExist(err) {
    fmt.Println("Error: Unable to find \"run.json\" in the current directory.")
    os.Exit(1)
  }

  // Read "run.json"
  data, err := ioutil.ReadFile("run.json")
  if err != nil {
    fmt.Println("Error: Unable to read \"run.json\".")
    os.Exit(1)
  }

  // Parse "run.json"
  commands := make(map[string]string)
  if err := json.Unmarshal([]byte(data), &commands); err != nil {
    fmt.Println("Error: Unable to parse \"run.json\".")
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

  // Check if command exists in "run.json"
  if _, ok := commands[command]; !ok {
    fmt.Printf("Error: Command \"%s\" cannot be found in \"run.json\".\n", command)
    os.Exit(2)
  }

  // Execute command
  cmd := exec.Command("/bin/sh", "-c", commands[command])
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    fmt.Printf("Error: Failed to execute \"%s\".\n", commands[command])
    os.Exit(1)
  }
}
