package commands

import (
  "fmt"
  "os"

  "github.com/lukecjohnson/run/lib/utils"
)

func List() {
  commands := utils.ParseYaml("run.yaml")

  fmt.Println("Available commands:")
  for command := range commands {
    fmt.Println("  " + command)
  }

  os.Exit(0)
}
