package commands

import (
  "fmt"
  "io/ioutil"
  "os"

  "gopkg.in/yaml.v3"
)

func List() {
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

  fmt.Println("\n" + "Available commands:")
  for command := range commands {
    fmt.Println("  " + command)
  }
  fmt.Println()

  os.Exit(0)
}
