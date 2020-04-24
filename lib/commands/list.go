package commands

import (
  "fmt"
  "os"

  "github.com/lukecjohnson/dot/lib/utils"
)

func List() {
  commands := utils.ParseYaml("dot.yaml")

  fmt.Println("\n" + "Available commands:")
  for command := range commands {
    fmt.Println("  " + command)
  }
  fmt.Println()

  os.Exit(0)
}
