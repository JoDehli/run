package commands

import (
  "fmt"
  "os"

  "github.com/lukecjohnson/run/lib/utils"
)

func Version() {
  currentVersion := utils.CurrentVersion
  fmt.Printf("run %v \n", currentVersion)

  os.Exit(0)
}
