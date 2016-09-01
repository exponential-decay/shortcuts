#shortcuts

A tool for extracting information from Windows Shortcut Files. 
Implemented in Golang.

##Standing On The Shoulder Of Giants

There's a whole host of information out there about this. I simply
wanted a library I could work with in some of my more recent code. 

Here are some links you may find useful. 

- Official Microsoft Specification: https://msdn.microsoft.com/en-us/library/dd871305.aspx 
- The Meaning of Life: http://computerforensics.parsonage.co.uk/downloads/TheMeaningofLIFE.pdf 
- ForensicWiki.org: http://www.forensicswiki.org/wiki/LNK 
- Just Solve It Wiki: http://fileformats.archiveteam.org/wiki/Windows_Shortcut 
- Windows LNK Parsing Utility: https://tzworks.net/prototype_page.php?proto_id=11 
- GitHub scan: https://github.com/search?utf8=%E2%9C%93&q=windows+shortcut

#Microsoft Shell Format

It is undocumented, but a good project i discovered reverse engineering it is: 
https://msdn.microsoft.com/en-us/library/windows/desktop/bb762540(v=vs.85).aspx

For C developers, this library looks invaluable.

##License

[GPLv3](https://github.com/exponential-decay/shortcutz/blob/master/LICENSE)

---

###NOTES ON STRUCTURES

Placing these here until we have better understanding. 

These structures seem to fall under two names:

* MS Shell Data-Source
* MS Shell Bag

0x31 Type Structures:

ITEM SIZE:        ## ##
TYPE:             31 00
SIZE:             00 00 00 00
DATETIME:         0C 49 0E 0A 
SIZE:             10 00 
8BIT NAME:        53 6F 75 72 63 65 00 00 
EXTSIZE:          3A 00 
VERSION:          08 00 
SIGNATURE:        04 00 EF BE 
DATE1:            2D 47 3B B2 
DATE2:            0C 49 0E 0A 
IDENTIFIER:       2A 00 
UNKNOWN:          00 00 
MFTENTRY:         9D AA 02 00 
MFTSEQ:           00 00 22 00 
UTF16 STRING:     00 00 00 00 00 00 00 00 00 00 00 00 00 00 53 00 6F 00 75 00 72 00 63 00 65 00 00 00 16 00


ITEM SIZE:        ## ##
ITEM TYPE:        31 00
SIZE:             00 00 00 00 
DATETIME:         09 49 0a 0d 
SIZE:             10 00 
8BIT NAME:        67 69 74 68 75 62 2e 63 6f 6d 00 00 
EXTSIZE:          42 00 
VERSION:          08 00 
SIGNATURE:        04 00 ef be 
DATE1:            14 47 53 39 
DATE2:            09 49 0a 0d 
IDENTIFIER:       2a 00 
UNKNOWN:          00 00 
MFTENTRY:         63 77 02 00 
MFTSEQ:           00 00 03 00 
UTF16 STRING:     00 00 00 00 00 00 00 00 00 00 00 00 00 00 67 00 69 00 74 00 68 00 75 00 62 00 2e 00 63 00 6f 00 6d 00 00 00 1a 00