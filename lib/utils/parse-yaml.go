package utils

import (
  "fmt"
  "io/ioutil"
  "os"

  "gopkg.in/yaml.v3"
)

func ParseYaml(filename string) map[string]string {
  data, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Printf("Error: Cannot find \"%s\" in the current directory.", filename)
    os.Exit(1)
  }

  yamlMap := make(map[string]string)
  if err := yaml.Unmarshal([]byte(data), &yamlMap); err != nil {
    fmt.Printf("Error: Unable to parse \"%s\".", filename)
    os.Exit(1)
  }

  return yamlMap
}