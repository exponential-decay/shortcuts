package main

//a way to separate what Microsoft have told us and what we've
//reverse engineered in the community. 

//Credit to all this work is in this repo: https://github.com/libyal/libfwsi
//@JoachimMetz

//Unknown: see https://msdn.microsoft.com/en-us/library/windows/desktop/cc144090(v=vs.85).aspx for clues...
var found_shortcut_PIDL = [17]byte{0x50, 0xe0, 0x4f, 0xd0, 0x20, 0xea, 0x3a, 0x69, 0x10, 0xa2, 0xd8, 0x08, 0x00, 0x2b, 0x30, 0x30, 0x9d}

//shell bag {search term in google also}
//https://files.sans.org/summit/Digital_Forensics_and_Incident_Response_Summit_2015/PDFs/PlumbingtheDepthsShellBagsEricZimmerman.pdf
//http://www.williballenthin.com/forensics/shellbags/

const beef = 0xbeef0004
const beeflen = 0x04
const beefseek = 0x08

//structs that make up the shortcut specification [76 bytes] 
type SHITEM_NTFS struct {
   ItemType             uint16
   Size1                uint32
   DosModifiedDate      uint16    //nb. these two may need swapping
   DosModifiedTime      uint16
   Size2                uint16
   //byte8string []byte    //8-bit string...
}

var stringpos8bit = 0x0C //12bytes

var SHITEM_NTFS_LEN = 0xC     //14 without 8-bit string
var EXT_LEN = 0x1C       //28bytes no []byte UTF16 block

var bit8string string
var utf16string string

type SHITEM_EXT_NTFS struct {
   Extsize        uint16
   Version        uint16
   Signature      uint32   //0xbeef0004
   CreatedDate    uint16   //creation
   CreatedTime    uint16
   ModifiedDate   uint16   //last accessed
   ModifiedTime   uint16
   Identifier     uint16   //maybe uint32 given 00 padding
   Unknown        uint16   //could be a uint32 in combination with identifier
   Mftentry       uint32   //e.g. 8c 75 06 00 == 0x0006758c
   Mftseqno       uint32   //e.g. 00 00 0A 00 == 10 ? 
   //utfstring   []byte   //what's left...
}

//values for bitwise in LinkFlags
//from: https://github.com/libyal/libfwsi/blob/master/documentation/Windows%20Shell%20Item%20format.asciidoc
var IdentifierFlagsMap = map[uint16]string{
   0x0:        nomapvalue,
   0x14:       "Windows XP or 2003",
   0x26:       "Windows Vista (SP0)",
   0x2a:       "Windows 2008, 7, 8.0",       //looks accurate to samples I have  
   0x2e:       "Windows 8.1, 10",
}

var ExtensionVersion = map[uint16]string{
   0x0:        nomapvalue,
   0x3:        "Windows XP or 2003",
   0x7:        "Windows Vista (SP0)",
   0x8:        "Windows 2008, 7, 8.0",       //looks accurate to samples I have  
   0x9:        "Windows 8.1, 10",
}
