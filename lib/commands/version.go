package commands

import (
  "fmt"
  "os"

  "github.com/lukecjohnson/dot/lib/utils"
)

func Version() {
  currentVersion := utils.CurrentVersion
  fmt.Println(currentVersion)

  os.Exit(0)
}
