package main

import (
   "os"
   "fmt"
   "bytes"
   "errors"
   "strings"
   "unicode/utf16"
   "encoding/binary"
)

type idlist struct {
   idlistlen   uint16
   idlistdata  []byte
   termid      uint16
}

type itemid struct {
   idsize uint16

   //shellitem data can be split like this...
   classid uint8
   data []byte
}

const uint16len = 2

func Extend(slice []uint16, element uint16) []uint16 {
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
         string_arr = Extend(string_arr, char)
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

func populateSHITEM_NTFS(class uint8, itemdata []byte, size uint16) {

   //SHITEM_NTFS
   //SHITEM_EXT_NTFS

   //fmt.Fprintf(os.Stderr, "data %x\n\n", itemdata)

   var t1 SHITEM_NTFS 
   var t2 SHITEM_EXT_NTFS

   if class == 0x1f {
      fmt.Fprintf(os.Stderr, "Computer: %s\n", string(itemdata[1:]))
   }

   if class == 0x2f {
      fmt.Fprintf(os.Stderr, "Drive: %s\n", string(itemdata[1:]))
   }

   if class >= 0x30 {

      bytereader := bytes.NewReader(itemdata[stringpos8bit:])
      var strpos = bytereader.Len()                   //length at time we begin...
      var eightbitstrlen int                          //length of the 8-bit string
      for bytereader.Len() > 0 {       
         val, _ := getint32(bytereader)               //we're looking for 0xbeef0004
         if val == beef {
            eightbitstrlen = (strpos - bytereader.Len()) - beefseek
            strpos = eightbitstrlen + stringpos8bit
            break
         }
         bytereader.Seek(-(beeflen-1), os.SEEK_CUR)   //beeflen minus one     //replace os.SEEK_CUR with io.SEEK...
      }

      bit8buf := itemdata[stringpos8bit:strpos-2]
      bit8string = string(bit8buf)

      pos := strpos + EXT_LEN
      remaining := len(itemdata)-(strpos + EXT_LEN) - 2     //lenght of uint16

      utf16buf := itemdata[pos:pos+remaining]
      utf16string = string(utf16buf)
      
      if len(utf16buf) % 2 == 0 {
         utf16stringdecoded, err := decodeUtf16(utf16buf)
         if err != nil {
            fmt.Println("Error decoding UTF-8 string in link target.")
            fmt.Println("8-bit name: ", bit8string)
         }
         fmt.Fprintf(os.Stdout, "8-bit: %s, UTF-16: %s\n", bit8string, utf16stringdecoded)
      }


      s1 := bytes.NewReader(itemdata[:stringpos8bit])
      s2 := bytes.NewReader(itemdata[strpos:])

      err := binary.Read(s1, binary.LittleEndian, &t1)
      if err != nil {
         fmt.Println(err)     //handle error
      }

      err = binary.Read(s2, binary.LittleEndian, &t2)
      if err != nil {
         fmt.Println(err)     //handle error
      }

   }
}

func getidfields(ids idlist) error {
   var start int
   var item itemid

   bytereader := bytes.NewReader(ids.idlistdata)

   for bytereader.Len() > 0 {
      sz, err := getint16(bytereader)
      item.idsize = sz
      if err != nil {
         return err
      }

      if item.idsize > 0 {
         curr_pos := bytereader.Size() - int64(bytereader.Len())
         item.classid, err = bytereader.ReadByte()
         if err != nil {
            return err
         }       

         //as there is no Peek() for a "bytes/reader" we can reset
         //like this... allowing the full struct to be output for debug
         bytereader.Seek(curr_pos, 0)

         readlen := (item.idsize-uint16len)     //minus one byte
         item.data = make([]byte, readlen)
         _, err = bytereader.Read(item.data[start:])
         if err != nil {   //likely io.EOF if we're not careful
            return err
         }
         populateSHITEM_NTFS(item.classid, item.data, item.idsize)
      }
   }
   return nil
}

func linktargethandler(fp *os.File, maskval uint32) error {
   var start int

   //link target structure handler
   if maskval == 0x1 {    
      var newIDList idlist

      intlen, err := fpgetlenint16(fp)
      if err != nil {
         return err
      }

      newIDList.idlistlen = intlen
      newIDList.idlistdata = make([]byte, newIDList.idlistlen)
      _, err = fp.Read(newIDList.idlistdata[start:])
      if err != nil {
         return err
      }

      //check for TerminalID = [0x00 0x00]
      if newIDList.termid == 0 {
         err = getidfields(newIDList)
         if err != nil {
            return err
         }
      } else {
         return errors.New("We haven't a TerminalID field equal to zero.")
      }
   }

   //haslinkinfo structure...
   if maskval == 0x2 {  

   }

   return nil
}
