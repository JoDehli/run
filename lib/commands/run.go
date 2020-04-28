package commands

import (
  "fmt"
  "os"
  "os/exec"

  "github.com/lukecjohnson/run/lib/utils"
)

func Run(command string) {
  commands := utils.ParseYaml("run.yaml")

  if _, ok := commands[command]; !ok {
    fmt.Printf("Error: Command \"%s\" cannot be found in \"run.yaml\".\n", command)
    os.Exit(2)
  }

  cmd := exec.Command("/bin/sh", "-c", commands[command])
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    fmt.Printf("Error: Failed to execute \"%s\".\n", commands[command])
    os.Exit(1)
  }
}
