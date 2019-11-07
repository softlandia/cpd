# code page detect #

(c) softlandia@gmail.com

golang library for detecting code page of text files  
support russian code page:

1. ASCII - default value
2. Windows1251
3. IBM866
4. KOI8R
5. UTF16LE only with bom
6. UTF16BE only with bom
7. UTF8

>download: go get -u github.com/softlandia/cpd  
>install: go install

## dependences ##

>"golang.org/x/text/encoding/charmap"  
>"golang.org/x/text/transform"

## types ##

IDCodePage uint16 - index of code page, support String() interface, you can fmt.Printf("code page index, name: %d, %s\n", cp, cp) where var cp received from cpd functions

## functions ##

1. CodePageDetect(r io.Reader, stopStr ...string) (IDCodePage, error)
2. FileCodePageDetect(fn string, stopStr ...string) (IDCodePage, error)
3. StrConvertCodePage(s string, fromCP, toCP uint16) (string, error)
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

    func StrConvertCodePage(s string, fromCP, toCP IDCodePage) (string, error)  //convert string from one code page to another

    func FileConvertCodePage(fileName string, fromCP, toCP IDCodePage) error    //convert code page file with "fileName"

## tests ##

coverage 96.2%  
folder "test_files" contain files for testing, do not remove/change/add if want support tests is work
