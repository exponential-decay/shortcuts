package main

import (
   "os"
   "fmt"
   "reflect"
)

var (
   headersz uintptr
)

func init() {
   structsizes()
}

func structsizes() {
   //first struct - shortcut header
   var header ShellLinkHeader
   headersz = reflect.TypeOf(header).Size()
}

//return: found, off1, off2, errors
func handleFile(fp *os.File) {

   //func Read(r io.Reader, order ByteOrder, data interface{}) error
   //buf := make([]byte, bfsize)
   ///=dataread, err := fp.Read(buf[start:])
   //check(err)

}

//callback for walk needs to match the following:
//type WalkFunc func(path string, info os.FileInfo, err error) error
func readFile (path string, fi os.FileInfo, err error) error {
   
   f, err := os.Open(path)
   defer f.Close()   //closing the file
   check(err)

   switch mode := fi.Mode(); {
   case mode.IsRegular():
      handleFile(f)
   case mode.IsDir():
      fmt.Fprintln(os.Stderr, "INFO:", fi.Name(), "is a directory.")      
   default: 
      fmt.Fprintln(os.Stderr, "INFO: Something completely different.")
   }
   return nil
}