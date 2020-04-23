package commands

import (
  "fmt"
  "os"
)

func Help() {
  usage := `
Usage:
  dot <command>        Runs the provided command found in "dot.yaml".

Other options:
  -h, --help           Prints this usage information.
  -l, --list           Prints all the available commands found in "dot.yaml".
  -v, --version        Prints the current installed version of this tool.

`
  fmt.Print(usage)

  os.Exit(0)
}
