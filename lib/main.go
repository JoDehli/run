package main

import (
  "flag"
  "fmt"
  "os"

  "github.com/lukecjohnson/dot/lib/commands"
)

func main() {
  arguments := os.Args[1:]

  versionFlag := flag.Bool("version", false, "Reports the current installed version of dot.")
  listFlag := flag.Bool("list", false, "Lists all the available commands found in \"dot.yaml\".")

  flag.Parse()

  flag.Usage = commands.Help

  if len(arguments) > 1 {
    fmt.Println("Error: Too many arguments.")
    fmt.Println("Run \"dot --help\" for usage instructions.")
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
