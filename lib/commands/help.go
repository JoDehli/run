package commands

import (
  "fmt"
  "os"
)

func Help() {

  // Usage
  usage := `
Usage:
  dot <script>      Runs the provided script found in "scripts.yaml".

Other options:
  -h, --help        Prints this usage information.
  -s, --scripts     Prints all the available scripts found in "scripts.yaml".
  -v, --version     Prints the current installed version of this tool.

`
  // Print usage
  fmt.Print(usage)

  // Exit
  os.Exit(0)

}
