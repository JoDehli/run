package commands

import (
  "fmt"

  "github.com/lukecjohnson/dot/lib/utils"
)

func Version() {
  currentVersion := utils.CurrentVersion
  fmt.Println(currentVersion)
}
