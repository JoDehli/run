package main

import (
  "flag"
  "fmt"
  "os"

  "github.com/lukecjohnson/run/lib/commands"
)

func main() {
  arguments := os.Args[1:]

  versionFlag := flag.Bool("version", false, "Reports the current installed version of run.")
  listFlag := flag.Bool("list", false, "Lists all the available commands found in \"run.yaml\".")

  flag.Usage = commands.Help

  flag.Parse()

  if len(arguments) > 1 {
    fmt.Println("Error: Too many arguments.")
    fmt.Println("Run \"run --help\" for usage instructions.")
    os.Exit(2)
  }

  if len(arguments) == 0 {
    commands.Help()
  }

  if *versionFlag {
    commands.Version()
  }

  if *listFlag {
    commands.List()
  }

  command := arguments[0]
  commands.Run(command)
}
