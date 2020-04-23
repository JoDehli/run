package commands

import (
  "fmt"
  "io/ioutil"
  "os"

  "gopkg.in/yaml.v3"
)

func Scripts() {

  // Read scripts.yaml
  data, err := ioutil.ReadFile("scripts.yaml")
  if err != nil {
    fmt.Println("Error: Cannot find scripts.yaml in the current directory.")
    os.Exit(1)
  }

  // Parse scripts.yaml
  scripts := make(map[string]string)
  if err := yaml.Unmarshal([]byte(data), &scripts); err != nil {
    fmt.Println("Error: Unable to parse scripts.yaml")
    os.Exit(1)
  }

  // Print available scripts
  fmt.Println()
  fmt.Println("Available scripts:")
  for script := range scripts {
    fmt.Printf("  %s \n", script)
  }
  fmt.Println()

  // Exit
  os.Exit(0)

}
