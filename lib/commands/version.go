package commands

import (
  "fmt"
  "os"

  "github.com/lukecjohnson/dot/lib/utils"
)

func Version() {
  currentVersion := utils.CurrentVersion
  fmt.Printf("dot %v \n", currentVersion)

  os.Exit(0)
}
