package main

import (
   "os"
   "fmt"
   "bytes"
   "errors"
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

func checklnkheader(HeaderSize [4]byte, ClassID [16]byte) bool {
   if header.HeaderSize != headsize || header.ClassID != classid {
      return false
   }
   return true
}

func readlinkflags(flags uint32) {
   var test uint32
   for i := 0; i < 32; i++ {
      if test == 0 {
         test = 1
      } else {
         test = test << 1         
      }

      value := LinkFlagsMap[flags & test]
      if value != nomapvalue {
         fmt.Println(value, "is set.")
      }
   }

}

//return: found, off1, off2, errors
func handleFile(fp *os.File) error {
   fmt.Println("\n")
   var start int
   buf := make([]byte, headersize)
   _, err := fp.Read(buf[start:])   
   check(err)

   b := bytes.NewReader(buf)

   //func Read(r io.Reader, order ByteOrder, data interface{}) error
   err = binary.Read(b, binary.LittleEndian, &header)
   if err != nil {
      return err
   }

   if !checklnkheader(header.HeaderSize, header.ClassID) {
      return errors.New("Not a valid shortcut file.")  //not a shortcut file... don't have to worry about doing too much
   }

   readlinkflags(header.LinkFlags)   
   return nil
}

//callback for walk needs to match the following:
//type WalkFunc func(path string, info os.FileInfo, err error) error
func readFile (path string, fi os.FileInfo, err error) error {
   
   f, err := os.Open(path)
   defer f.Close()   //closing the file
   check(err)

   switch mode := fi.Mode(); {
   case mode.IsRegular():
      err := handleFile(f)
      if err != nil {
         fmt.Fprintln(os.Stderr, "INFO:", fi.Name(), err)
      }
   case mode.IsDir():
      fmt.Fprintln(os.Stderr, "INFO:", fi.Name(), "is a directory.")      
   default: 
      fmt.Fprintln(os.Stderr, "INFO: Something completely different.")
   }
   return nil
}