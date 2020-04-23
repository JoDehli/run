package commands

import (
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"

  "gopkg.in/yaml.v3"
)

func Run(script string) {
  data, err := ioutil.ReadFile("scripts.yaml")
  if err != nil {
    fmt.Println("Error: Cannot find \"scripts.yaml\" in the current directory.")
    os.Exit(1)
  }

  scripts := make(map[string]string)
  if err := yaml.Unmarshal([]byte(data), &scripts); err != nil {
    fmt.Println("Error: Unable to parse \"scripts.yaml\".")
    os.Exit(1)
  }

  if _, ok := scripts[script]; !ok {
    fmt.Printf("Error: Script \"%s\" cannot be found in \"scripts.yaml\".\n", script)
    os.Exit(2)
  }

  cmd := exec.Command("/bin/sh", "-c", scripts[script])
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    fmt.Printf("Error: Failed to execute \"%s\".\n", scripts[script])
    os.Exit(1)
  }
}
