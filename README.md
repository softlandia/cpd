# code page detect #

(c) softlandia@gmail.com

golang 

>download: go get -u github.com/softlandia/cpd  
>install: go install

## dependences ##

>"golang.org/x/text/encoding/charmap"  
>"golang.org/x/text/transform"

## functions ##

1. StrConvertCodePage(s string, fromCP, toCP uint16) (string, error)
2. FileConvertCodePage(fileName string, fromCP, toCP uint16) error
3. FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error)
4. CodePageDetect(fn string) (int, error)

## description ##


    func StrConvertCodePage(s string, fromCP, toCP int64) (string, error)  //convert string from one code page to another

    func FileConvertCodePage(fileName string, fromCP, toCP int64) error //convert code page test file

    func FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error)  //search in path files with extention == fileNameExt and put file name to slice fileList

    func CodePageDetect(fn string, stopStr ...string) (int, error)  
    detect code page of text file "fn", 
    detect only IBM CodePage866 and Windows1251  
    return constant cpd.CpIBM866, cpd.CpWindows1251, cpd.CpASCII
    if string stopStr is present then input file scanned befor appearance stopStr

## tests ##

coverage 96.2%  
folder "test_files" contain files for testing, no remove/change/add
