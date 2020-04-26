package commands

import (
  "fmt"
  "os"

  "github.com/lukecjohnson/dot/lib/utils"
)

func List() {
  commands := utils.ParseYaml("dot.yaml")

  fmt.Println("Available commands:")
  for command := range commands {
    fmt.Println("  " + command)
  }

  os.Exit(0)
}
