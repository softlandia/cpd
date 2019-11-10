# code page detect #

(c) softlandia@gmail.com

>download: go get -u github.com/softlandia/cpd  
>install: go install

golang library for detecting code page of text files  
support russian code page:

1. ASCII - default value
2. Windows1251
3. IBM866
4. KOI8R
5. UTF-16le
6. UTF-16le
7. UTF-8
8. ISO8859-5

### feature ###
if file contain only latin symbols, this file detected as UTF-8  
this is not a mistake, this is a completely correct statement

on go vertion 1.12.6 add to exe 240 kB

## dependences ##

>"golang.org/x/text/encoding/charmap"  
>"golang.org/x/text/transform"

## types ##

IDCodePage uint16 - index of code page, support String() interface, you can fmt.Printf("code page index, name: %d, %s\n", cp, cp) where var cp received from cpd functions

## functions ##

1. CodePageDetect(r io.Reader, stopStr ...string) (IDCodePage, error)
2. FileCodePageDetect(fn string, stopStr ...string) (IDCodePage, error)
3. StrConvertCodePage(s string, fromCP, toCP IDCodePage) (string, error)
4. FileConvertCodePage(fileName string, fromCP, toCP IDCodePage) error

## description ##

    CodePageDetect(r io.Reader, stopStr ...string) (IDCodePage, error)
      detect code page of ascii data from reader 'r' 

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

coverage: 84.0% of statements  
folder "test_files" contain files for testing, do not remove/change/add if want support tests is work
