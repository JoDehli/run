package utils

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

func ParseJson(filename string) map[string]string {
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    fmt.Printf("Error: Unable to find \"%s\" in the current directory. \n", filename)
    os.Exit(1)
  }

  data, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Printf("Error: Unable to read \"%s\". \n", filename)
    os.Exit(1)
  }

  jsonMap := make(map[string]string)
  if err := json.Unmarshal([]byte(data), &jsonMap); err != nil {
    fmt.Printf("Error: Unable to parse \"%s\". \n", filename)
    os.Exit(1)
  }

  return jsonMap
}