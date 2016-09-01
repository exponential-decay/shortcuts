package main

import (
   "os"
   "fmt"
   "bytes"
   "errors"
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

   fmt.Fprintf(os.Stderr, "data %x\n\n", itemdata)

   var test SHITEM_NTFS 
   test.itemsize = size

   if class >= 0x30 {
      bytereader := bytes.NewReader(itemdata[stringpos8bit:])
      var strpos = bytereader.Len()
      var strlen int
      for bytereader.Len() > 0 {
         val, _ := getint32(bytereader)
         if val == beef {
            strlen = (strpos - bytereader.Len()) - beefseek
            strpos = strlen + stringpos8bit
            break
         }
         bytereader.Seek(-(beeflen-1), os.SEEK_CUR)
      }
      eightbitstring := string(itemdata[stringpos8bit:strpos])
      fmt.Println(eightbitstring)
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
