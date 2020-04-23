package commands

import (
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"

  "gopkg.in/yaml.v3"
)

func Run(command string) {
  data, err := ioutil.ReadFile("dot.yaml")
  if err != nil {
    fmt.Println("Error: Cannot find \"dot.yaml\" in the current directory.")
    os.Exit(1)
  }

  commands := make(map[string]string)
  if err := yaml.Unmarshal([]byte(data), &commands); err != nil {
    fmt.Println("Error: Unable to parse \"dot.yaml\".")
    os.Exit(1)
  }

  if _, ok := commands[command]; !ok {
    fmt.Printf("Error: Command \"%s\" cannot be found in \"dot.yaml\".\n", command)
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
