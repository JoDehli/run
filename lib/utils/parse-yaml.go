package utils

import (
  "fmt"
  "io/ioutil"
  "os"

  "gopkg.in/yaml.v3"
)

func ParseYaml(filename string) map[string]string {
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    fmt.Printf("Error: Unable to find \"%s\" in the current directory. \n", filename)
    os.Exit(1)
  }

  data, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Printf("Error: Unable to read \"%s\". \n", filename)
    os.Exit(1)
  }

  yamlMap := make(map[string]string)
  if err := yaml.Unmarshal([]byte(data), &yamlMap); err != nil {
    fmt.Printf("Error: Unable to parse \"%s\". \n", filename)
    os.Exit(1)
  }

  return yamlMap
}