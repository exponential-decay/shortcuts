package main

import (
   "os"
   "time"
   "bytes"
   "strings"
   "unicode/utf16"
   "encoding/binary"   
)

const uint16len = 2

func ExtendSlice(slice []uint16, element uint16) []uint16 {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]uint16, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}

func decodeUtf16(utf16buf []byte) (string, error) {
   const SHORT_LEN = 2
   var x = 0
   var string_arr []uint16   
   for x < len(utf16buf) {
      tmpbuf := utf16buf[x:2+x]
      var char uint16
      strbuf := bytes.NewReader(tmpbuf)
      err := binary.Read(strbuf, binary.LittleEndian, &char)
      if err != nil {
         return "", err
      }
      if char != 0x00 {
         string_arr = ExtendSlice(string_arr, char)
      }
      x+=2
   }
   utf16decoded := strings.TrimSpace(string(utf16.Decode(string_arr)))
   return utf16decoded, nil
}

func getint32(bytereader *bytes.Reader) (uint32, error) {
   var newint uint32
   err := binary.Read(bytereader, binary.LittleEndian, &newint)
   if err != nil {
      return 0, err
   }
   return newint, err
}

func getint16(bytereader *bytes.Reader) (uint16, error) {
   var newint uint16

   //func Read(r io.Reader, order ByteOrder, data interface{}) error
   err := binary.Read(bytereader, binary.LittleEndian, &newint)
   if err != nil {
      return 0, err
   }
   return newint, err
}

//get uint16 from beginning of an appropriate file stream
func fpgetlenint16(fp *os.File) (uint16, error) {
   var start int
   buf := make([]byte, uint16len)
   _, err := fp.Read(buf[start:])
   check(err)
   newint, err := getint16(bytes.NewReader(buf))
   return newint, err
}

//from https://golang.org/src/archive/zip/struct.go
func msDosTimeToTime(dosDate, dosTime uint16) time.Time {
   return time.Date(
      //date bits 0-4: day of month; 5-8: month; 9-15: years since 1980
      int(dosDate>>9+1980),
      time.Month(dosDate>>5&0xf),
      int(dosDate&0x1f),

      //time bits 0-4: second/2; 5-10: minute; 11-15: hour
      int(dosTime>>11),
      int(dosTime>>5&0x3f),
      int(dosTime&0x1f*2),
      0, // nanoseconds

      time.UTC,
   )
}