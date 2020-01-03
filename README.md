# code page detect #

(c) softlandia@gmail.com

>download: go get -u github.com/softlandia/cpd  
>install: go install

golang library for detecting code page of text files  
multibyte code pages and single-byte Russian code pages are supported:

no ID                Name		uint16

1. ASCII:            "ASCII",		3
2. ISOLatinCyrillic: "ISO-8859-5",	8
3. CP866:            "CP866",		2086
4. Windows1251:      "Windows-1251",	2251
5. UTF8:             "UTF-8",		106
6. UTF16LE:          "UTF-16LE",	1014
7. UTF16BE:          "UTF-16BE",	1013
8. KOI8R:            "KOI8-R",		2084
9. UTF32LE:          "UTF-32LE",	1019
10. UTF32BE:         "UTF-32BE",	1018


## feature ##

encoding is determined both by the presence of the bom attribute and by heuristic  
if file contain only latin symbols from first half of code page, this file detected as UTF-8  
this is not a mistake, this is a completely correct statement

on go vertion 1.12.6 add to exe 250 kB

## dependences ##

>"golang.org/x/text/encoding/charmap"  
>"golang.org/x/text/transform"

## types ##

IDCodePage uint16 - index of code page, support String() interface, you can fmt.Printf("code page index, name: %d, %s\n", cp, cp) where var cp received from cpd functions

## variables ##

ReadBufSize int = 1024 // count of byte to read from input reader by default

## functions ##

1. CodePageDetect(r io.Reader, stopStr ...string) (IDCodePage, error)
2. FileCodePageDetect(fn string, stopStr ...string) (IDCodePage, error)

## description ##

    func CodePageAutoDetect(content []byte) (result IDCodePage) 
      autodetect code page from input slice of byte
      use this function instead golang.org/x/net/html/charset.DetermineEncoding()

    CodePageDetect(r io.Reader, stopStr ...string) (IDCodePage, error)
      detect code page of ascii data from reader 'r' 
      use library 'reflect' to check input reader
      default read only first 1024 byte from 'r' (var ReadBufSize to change this setting)
      input parameter stopStr not using

    FileCodePageDetect(fn string, stopStr ...string) (IDCodePage, error)
      detect code page of text file "fn", read first 1024 byte (var ReadBufSize to change this setting)
      return error if problem with file "fn"
      return cpd.ASCII if code page not detected
      return one of next constant (code_pages_id.go): cpd.IBM866, cpd.Windows1251, cpd.KOI8R, cpd.UTF8, UTF16LE, UTF16BE
      file must contain characters of the Rusian alphabet
      string stopStr now not using

    func StrConvertCodePage(s string, fromCP, toCP IDCodePage) (string, error)  //convert string from one code page to another, support Windows1251 & IBM866

    func FileConvertCodePage(fileName string, fromCP, toCP IDCodePage) error    //convert code page file with "fileName", support Windows1251 & IBM866

## tests ##

coverage: 79% of statements  
folder "test_files" contain files for testing, do not remove/change/add if want support tests is work
