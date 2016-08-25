//https://msdn.microsoft.com/en-us/library/dd871305.aspx
package main 

//things we know that we're looking for in the shortcut file
var headsize = [4]byte{0x4C, 0x00, 0x00, 0x00}
var classid = [16]byte{0x01, 0x14, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}

//structs that make up the shortcut specification
type ShellLinkHeader struct {
   HeaderSize  [4]byte           //HeaderSize
   ClassID     [16]byte          //LinkCLSID
   LinkFlags   [4]byte           //LinkFlags
   FileAttr    [4]byte           //FileAttributes
   Creation    [8]byte           //CreationTime
   Access      [8]byte           //AccessTime
   Write       [8]byte           //WriteTime
   FileSz      [4]byte           //FileSize
   IconIndex   [4]byte           //IconIndex
   ShowCmd     [4]byte           //ShowCommand
   HotKey      [2]byte           //HotKey
   Reserved1   [2]byte           //Reserved1
   Reserved2   [4]byte           //Reserved2
   Reserved3   [4]byte           //Reserved3
}