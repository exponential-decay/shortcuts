//https://msdn.microsoft.com/en-us/library/dd871305.aspx
package main 

//things we know that we're looking for in the shortcut file
var headsize = [4]byte{0x4C, 0x00, 0x00, 0x00}
var classid = [16]byte{0x01, 0x14, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}

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

//structs that make up the shortcut specification
type ShellLinkHeader struct {
   HeaderSize  [4]byte           //HeaderSize
   ClassID     [16]byte          //LinkCLSID
   LinkFlags   uint32            //LinkFlags
   //LinkFlags   [4]byte         //LinkFlags
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