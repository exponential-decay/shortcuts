package main

import (
   "os"
   "fmt"
   "bytes"
   "errors"
   "encoding/binary"
)

func formatGlobalID(gid []byte) string {
   //e.g. 208d2c60-3aea-1069-a2d7-08002b30309d
   //e.g. e04fd020ea3a6910a2d808002b30309d becomes e04fd020-ea3a-6910-a2d8-08002b30309d
   if len(gid) >= 18 {
      clid := fmt.Sprintf("%x-%x-%x-%x-%x\n", gid[2:2+4], gid[6:6+2], gid[8:8+2], gid[10:10+2], gid[12:])
      if clid != "" {
         if WindowsClassIDs[clid] != "" {
            return WindowsClassIDs[clid]
         } else {
            return clid
         }
      }
   }
   return ""
}

func populateSHITEM_NTFS(class uint8, itemdata []byte, size uint16) {
   var t1 SHITEM_NTFS 
   var t2 SHITEM_EXT_NTFS

   if class == 0x1f {
      fmt.Fprintf(os.Stderr, "Windows Class Identifier: %s\n", formatGlobalID(itemdata[:]))
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

      bit8buf := bytes.Trim(itemdata[stringpos8bit:strpos-1], "\x00")
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

      fmt.Println("DOS:", msDosTimeToTime(t1.DosModifiedDate, t1.DosModifiedTime))
      fmt.Println("NTFS:", msDosTimeToTime(t2.CreatedDate, t2.CreatedDate), msDosTimeToTime(t2.ModifiedDate, t2.ModifiedDate))      
      if ExtensionVersion[t2.Version] != IdentifierFlagsMap[t2.Identifier] {
         fmt.Fprintf(os.Stderr, "Extension Version: %s\n", ExtensionVersion[t2.Version])
         fmt.Fprintf(os.Stderr, "Identifier Value: %s\n", IdentifierFlagsMap[t2.Identifier])
      } else {
         fmt.Fprintf(os.Stderr, "Identifier Value: %s\n", IdentifierFlagsMap[t2.Identifier])
      }
      fmt.Println("---")
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
         //like this... also allows the full struct to be output for debug
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
