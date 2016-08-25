package main 

import (
   "os"
   "fmt"
)

func check(err error) {
   if err != nil {
      fmt.Fprintln(os.Stderr, "ERROR:", err)
   }
}