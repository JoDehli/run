package commands

import (
  "flag"
  "fmt"
  "os"
)

func Help() {
  fmt.Println("Usage:")
  fmt.Println("  run <command>" + "\t\t" + "Runs the provided command found in \"run.yaml\".")

  fmt.Println("\n" + "Other options:")
  flag.VisitAll(func(f *flag.Flag) {
    fmt.Println("  --" + f.Name + "\t\t" + f.Usage)
  })

  os.Exit(0)
}
