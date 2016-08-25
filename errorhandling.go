package main 

import (
   "os"
   "fmt"
)

func check(err error) {
   if err != nil {
      fmt.Fprintln(os.Stderr, "ERROR:", err)
      os.Exit(1)  //should only exit if root is null, consider no-exit
   }
}