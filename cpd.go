//Package cpd - code page detect
// (c) 2019 softlandia@gmail.com
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

	//is buf contains the BOM of utf-8, utf-16le or utf-16be
	if idCodePage, ok := CheckBOM(buf); ok {
		return idCodePage, nil
	}

	if ValidUTF8(buf) {
		return UTF8, nil
	}

	return CodePageAutoDetect(buf), nil
}

//CodePageAutoDetect - auto detect code page of input content
func CodePageAutoDetect(content []byte) (result IDCodePage) {
	return CodepageDic.Match(content) //TODO большинству матчеров требуется более 2х символов, надо проверить на минимальную длину
}

//FileConvertCodePage - replace code page text file from one to another
// support convert only from/to Windows1251/IBM866
func FileConvertCodePage(fileName string, fromCP, toCP IDCodePage) error {
	if fromCP == toCP {
		return nil
	}

	if (fromCP != Windows1251) && (fromCP != CP866) {
		return nil
	}

	if (toCP != Windows1251) && (toCP != CP866) {
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
			return fmt.Errorf("code page convert error on file '%s': %v", fileName, err)
		}
		fmt.Fprintf(oFile, "%s\n", s)
	}
	oFile.Close()
	iFile.Close()
	return os.Rename(tmpFileName, fileName)
}

//ToUTF8 -
//TODO need realization
func ToUTF8(s string) string {
	return s
}

//StrConvertCodePage - convert string from one code page to another
// function for future, at now support convert only from/to Windows1251/IBM866
func StrConvertCodePage(s string, fromCP, toCP IDCodePage) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if fromCP == toCP {
		return s, nil
	}

	var err error

	switch fromCP {
	case CP866:
		s, _, err = transform.String(charmap.CodePage866.NewDecoder(), s)
	case Windows1251:
		s, _, err = transform.String(charmap.Windows1251.NewDecoder(), s)
	}
	switch toCP {
	case CP866:
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
