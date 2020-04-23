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
  -h, --help        Print this usage information.
  -v, --version     Prints the current installed version of this tool.

`

  fmt.Print(usage)
  os.Exit(0)
}