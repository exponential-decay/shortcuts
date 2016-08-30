package main

//a way to separate what Microsoft have told us and what we've
//reverse engineered in the community. 

//Credit to all this work is in this repo: https://github.com/libyal/libfwsi
//@JoachimMetz

//Unknown: see https://msdn.microsoft.com/en-us/library/windows/desktop/cc144090(v=vs.85).aspx for clues...
var found_shortcut_PIDL = [17]byte{0x50, 0xe0, 0x4f, 0xd0, 0x20, 0xea, 0x3a, 0x69, 0x10, 0xa2, 0xd8, 0x08, 0x00, 0x2b, 0x30, 0x30, 0x9d}