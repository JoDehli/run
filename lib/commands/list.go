package commands

import (
  "fmt"
  "os"

  "github.com/lukecjohnson/run/lib/utils"
)

func List() {
  commands := utils.ParseJson("run.json")

  fmt.Println("Available commands:")
  for command := range commands {
    fmt.Println("  " + command)
  }

  os.Exit(0)
}
