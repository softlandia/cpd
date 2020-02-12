package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/softlandia/cpd"
)

//выводит кодировку всех файлов с указанным расширением
//пример запуска: >detect-all-files .txt
func main() {
	var fl []string
	FindFilesExt(&fl, ".\\", os.Args[1])
	for _, fn := range fl {
		t, _ := cpd.FileCodepageDetect(fn)
		fmt.Printf("file: \t`%s`\t`%s`\n", fn, t)
	}
}

//FindFilesExt - search all files in path with 'ext' & put to list
//path - "c:\tmp"
//ext  - ".log"
//sample:  n, err := FindFilesExt(&fl, "c:\\tmp", ".log")
func FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error) {
	if fileList == nil {
		return 0, errors.New("first parameter 'fileList' is nil")
	}
	extFile := strings.ToUpper(fileNameExt)
	i := 0 //index founded files
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			//skip folders
			return nil
		}
		if strings.ToUpper(filepath.Ext(path)) != extFile {
			//skip folders and files with extention not extFile
			return nil
		}
		//file found
		i++
		*fileList = append(*fileList, path)
		return nil
	})
	return i, err
}
