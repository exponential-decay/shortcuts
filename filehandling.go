package main

import (
   "os"
   "fmt"
   "bytes"
   "encoding/binary"
   "reflect"
)

var (
   header ShellLinkHeader
   headersize uintptr
)

func init() {
   structsizes()
}

func structsizes() {
   //first struct - shortcut header
   headersize = reflect.TypeOf(header).Size()
}

//return: found, off1, off2, errors
func handleFile(fp *os.File) {
   var start int
   buf := make([]byte, headersize)
   _, err := fp.Read(buf[start:])   
   check(err)

   b := bytes.NewReader(buf)

   //func Read(r io.Reader, order ByteOrder, data interface{}) error
   err = binary.Read(b, binary.LittleEndian, &header)
   check(err)

   fmt.Println(string(header.HeaderSize))
   fmt.Println(header.ClassID)
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