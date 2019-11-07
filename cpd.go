//Package cpd - code page detect
// (c) 2019 softlandia@gmail.com
// v0.1.0
// 01/oct/2019
package cpd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

//ReadBufSize - byte count for reading from file, func FileCodePageDetect()
var ReadBufSize int = 1024

//CodePageAutoDetect - auto detect code page of input content
func CodePageAutoDetect(content []byte) (result IDCodePage) {
	return CodePages.Match(content)
}

//CodePageDetect - detect code page of ascii data from reader 'r'
func CodePageDetect(r io.Reader, stopStr ...string) (IDCodePage, error) {
	//initial test
	//test input interfase
	if !reflect.ValueOf(r).IsValid() {
		return ASCII, fmt.Errorf("input reader is nil")
	}

	//make slice of byte from input reader
	buf, err := bufio.NewReader(r).Peek(ReadBufSize)
	if (err != nil) && (err.Error() != "EOF") {
		return ASCII, err
	}

	//check file header // utf-8, utf-16 with BOM
	if idHeader, ok := checkHeader(buf); ok {
		return idHeader, nil
	}
	return CodePageAutoDetect(buf), nil
}

//FileCodePageDetect - detect code page of text file
func FileCodePageDetect(fn string, stopStr ...string) (IDCodePage, error) {

	iFile, err := os.Open(fn)
	if err != nil {
		return ASCII, err
	}
	defer iFile.Close()

	if len(stopStr) > 0 {
		return CodePageDetect(iFile, stopStr[0])
	}
	return CodePageDetect(iFile)
}

//FileConvertCodePage - replace code page text file from one to another
func FileConvertCodePage(fileName string, fromCP, toCP IDCodePage) error {
	if fromCP == toCP {
		return nil
	}

	if (fromCP != Windows1251) && (fromCP != IBM866) {
		return nil
	}

	if (toCP != Windows1251) && (toCP != IBM866) {
		return nil
	}

	iFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer iFile.Close()

	//TODO need using sytem tmp folder
	tmpFileName := fileName + "~"
	oFile, err := os.Create(tmpFileName)
	if err != nil {
		return err
	}
	defer oFile.Close()

	s := ""
	iScanner := bufio.NewScanner(iFile)
	for i := 0; iScanner.Scan(); i++ {
		s = iScanner.Text()
		s, err = StrConvertCodePage(s, fromCP, toCP)
		if err != nil {
			oFile.Close()
			os.Remove(tmpFileName)
			return fmt.Errorf("cde page convert error on file '%s': %v", fileName, err)
		}
		fmt.Fprintf(oFile, "%s\n", s)
	}
	oFile.Close()
	iFile.Close()
	return os.Rename(tmpFileName, fileName)
}

//StrConvertCodePage - convert string from one code page to another
func StrConvertCodePage(s string, fromCP, toCP IDCodePage) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if fromCP == toCP {
		return s, nil
	}

	var err error

	switch fromCP {
	case IBM866:
		s, _, err = transform.String(charmap.CodePage866.NewDecoder(), s)
	case Windows1251:
		s, _, err = transform.String(charmap.Windows1251.NewDecoder(), s)
	}
	switch toCP {
	case IBM866:
		s, _, err = transform.String(charmap.CodePage866.NewEncoder(), s)
	case Windows1251:
		s, _, err = transform.String(charmap.Windows1251.NewEncoder(), s)
	}
	return s, err
}

// CodePageAsString - return name of char set with id codepage
// if codepage not exist - return ""
func CodePageAsString(codepage IDCodePage) string {
	return codePageName[codepage]
}
