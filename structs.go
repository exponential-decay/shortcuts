//https://msdn.microsoft.com/en-us/library/dd871305.aspx
package main 

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