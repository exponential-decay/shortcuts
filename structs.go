//https://msdn.microsoft.com/en-us/library/dd871305.aspx
package main 

//NB Other enumerators that we may need
//Shell format spec:
//https://github.com/libyal/libfwsi/blob/110cc8e0f2d549785cc4b1a2b08877a47f61e75a/documentation/Windows%20Shell%20Item%20format.asciidoc 
//Shell link data flags enum Msoft:
//https://msdn.microsoft.com/en-us/library/windows/desktop/bb762540(v=vs.85).aspx

//things we know that we're looking for in the shortcut file
var headsize = [4]byte{0x4C, 0x00, 0x00, 0x00}
var classid = [16]byte{0x01, 0x14, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}

//structs that make up the shortcut specification [76 bytes] 
type ShellLinkHeader struct {
   HeaderSize  [4]byte           //HeaderSize
   ClassID     [16]byte          //LinkCLSID
   LinkFlags   uint32            //LinkFlags      [4]byte
   FileAttr    uint32            //FileAttributes [4]byte
   Creation    [8]byte           //CreationTime
   Access      [8]byte           //AccessTime
   Write       [8]byte           //WriteTime
   FileSz      [4]byte           //FileSize
   IconIndex   [4]byte           //IconIndex
   ShowCmd     [4]byte           //ShowCommand

   //[2]byte HotKey values for shortcut shortcuts
   HotKeyLow   byte              //HotKeyLow
   HotKeyHigh  byte              //HotKeyHigh
   
   Reserved1   [2]byte           //Reserved1
   Reserved2   [4]byte           //Reserved2
   Reserved3   [4]byte           //Reserved3
}

//maps
var nomapvalue = "nil"

//values for bitwise in LinkFlags
var LinkFlagsMap = map[uint32]string{
   0x0: nomapvalue,
   0x1: "HasLinkTargetIDList",
   0x2: "HasLinkInfo",
   0x4: "HasName",         
   0x8: "HasRelativePath",
   0x10: "HasWorkingDir", 
   0x20: "HasArguments",
   0x40: "HasIconLocation",
   0x80: "IsUnicode",
   0x100: "FoceNoLinkInfo",
   0x200: "HasExpString",
   0x400: "RunInSeparateProcess", 
   0x800: "Unused1",
   0x1000: "HasDarwinID",
   0x2000: "RunAsUser",
   0x4000: "HasExpIcon",
   0x8000: "NoPidAlias",
   0x10000: "Unused2",
   0x20000: "RunWithShimLayer",
   0x40000: "ForceNoLinkTrack",
   0x80000: "EnableTargetMetadata",
   0x100000: "DisableLinkPathTracking",
   0x200000: "DisableKnownFolderTracking",
   0x400000: "DisableKnownFolderAlias",
   0x800000: "AllowLinkToLink",
   0x1000000: "UnaliasOnSave",
   0x2000000: "PreferEnvironmentPath", 
   0x4000000: "KeepLocalIDListForUNCTarget",
   0x8000000: "Unused3",
   0x10000000: "Unused4",
   0x20000000: "Unused5",
   0x40000000: "Unused6",
   0x80000000: "Unused7",
}

//values for bitwise in FileAttributes
var FileAttrMap = map[uint32]string{
   0x0: nomapvalue,
   0x1: "FILE ATTRIBUTE READ ONLY",
   0x2: "FILE ATTRIBUTE HIDDEN",
   0x4: "FILE ATTRIBUTE SYSTEM",         
   0x8: "Reserved 1",
   0x10: "FILE ATTRIBUTE DIRECTORY", 
   0x20: "FILE ATTRIBUTE ARCHIVE",
   0x40: "Reserved 2",
   0x80: "FILE ATTRIBUTE NORMAL",
   0x100: "FILE ATTRIBUTE TEMPORARY",
   0x200: "FILE ATTRIBUTE SPARSE FILE",
   0x400: "FILE ATTRIBUTE REPARSE POINT", 
   0x800: "FILE ATTRIBUTE COMPRESSED",
   0x1000: "FILE ATTRIBUTE OFFLINE",
   0x2000: "FILE ATTRIBUTE NOT CONTENT INDEXED",   
   0x4000: "FILE ATTRIBUTE ENCRYPTED",   
}

//Verbose, but accurate... HotKeyMapping
//LowByte represents a value, must be *one* of these
var HotKeyMapLow = map[byte]string{

   0x0: nomapvalue,
   //numbers
   0x30: "0",
   0x31: "1",
   0x32: "2",
   0x33: "3",
   0x34: "4",
   0x35: "5",
   0x36: "6",
   0x37: "7",
   0x38: "8",
   0x39: "9",

   //A-Z
   0x41: "A",
   0x42: "B",
   0x43: "C",
   0x44: "D",
   0x45: "E",
   0x46: "F",
   0x47: "G",
   0x48: "H",
   0x49: "I",
   0x4A: "J",
   0x4B: "K",
   0x4C: "L",
   0x4D: "M",
   0x4E: "N",
   0x4F: "O",
   0x50: "P",
   0x51: "Q",
   0x52: "R",
   0x53: "S",
   0x54: "T",
   0x55: "U",
   0x56: "V",
   0x57: "W",
   0x58: "X",
   0x59: "Y",
   0x5A: "Z",

   //Function Keys
   0x70: "F1",
   0x71: "F2",
   0x72: "F3",
   0x73: "F4",
   0x74: "F5",
   0x75: "F6",
   0x76: "F7",
   0x77: "F8",
   0x78: "F9",
   0x79: "F10",
   0x7A: "F11",
   0x7B: "F12",
   0x7C: "F13",
   0x7D: "F14",
   0x7E: "F15",
   0x7F: "F16",
   0x80: "F17",
   0x81: "F18",
   0x82: "F19",
   0x83: "F20",
   0x84: "F21",
   0x85: "F22",
   0x86: "F23",
   0x87: "F24",

   //control keys
   0x90: "NUM LOCK",
   0x91: "SCROLL LOCK",

}

//HotKeyMap High 
//HighByte represents a value, must be one or a combination of these
var HotKeyMapHigh = map[byte]string{
   0x0: nomapvalue,
   0x01: "SHIFT",
   0x02: "CTRL",
   0x04: "ALT",
}