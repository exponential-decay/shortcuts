package main

import (
   "fmt"
   "unsafe"
)

func main() {
   fmt.Println("4C0000000114020000000000C000000000000046")

   var ross ShellLinkHeader
   const infoSize = unsafe.Sizeof(ross)

   fmt.Println(infoSize)
}
   