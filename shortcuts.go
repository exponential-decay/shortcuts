package main

import (
   "os"
   "fmt"
   "flag"
   "path/filepath"   
)

var (
   version  string = "v0.0.1-beta"
   vers     bool
   file     string
)

func init() {
   flag.StringVar(&file, "file", "false", "File to find the distance between.")
   flag.BoolVar(&vers, "version", false, "[Optional] Return version of bindist.")   
}

func processfiles() {
   filepath.Walk(file, readFile)
}

func main() {

   flag.Parse()
   var verstring = "shortcuts version"
   if vers {
      fmt.Fprintf(os.Stderr, "%s %s \n", verstring, version)
      os.Exit(0)
   } else if flag.NFlag() <= 0 {    // can access args w/ len(os.Args[1:]) too
      fmt.Fprintln(os.Stderr, "Usage:  shortcutz [-file ...]")
      fmt.Fprintln(os.Stderr, "                  [Optional -version]")
      fmt.Fprintln(os.Stderr, "")
      fmt.Fprintln(os.Stderr, "Output: [STRING] TBD Structure of some sort...")
      fmt.Fprintf(os.Stderr, "Output: [STRING] '%s ...'\n\n", verstring)
      flag.Usage()
      os.Exit(0)
   }

   processfiles()
}
   