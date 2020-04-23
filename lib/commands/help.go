package commands

import (
  "fmt"
  "os"
)

func Help() {
  usage := `
Usage:
  dot <script>      Runs the provided script found in "scripts.yaml".

Other options:
  -h, --help        Prints this usage information.
  -s, --scripts     Prints all the available scripts found in "scripts.yaml".
  -v, --version     Prints the current installed version of this tool.

`
  fmt.Print(usage)

  os.Exit(0)
}
