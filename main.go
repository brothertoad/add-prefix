package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) < 3 {
    fmt.Printf("Need at least two arguments\n")
    os.Exit(1)
  }
  prefix := os.Args[1]
  // fmt.Printf("prefix is %s\n", prefix)
  files := os.Args[2:]
  for _, file := range(files) {
    target := prefix + file
    fmt.Printf("Moving %s to %s...\n", file, target)
    if err := os.Rename(file, target); err != nil {
      fmt.Printf("Unable to rename %s to %s: %s\n", file, target, err.Error())
    }
  }
}
