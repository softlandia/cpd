//Package cpd - code page detect
// (c) 2020 softlandia@gmail.com
package cpd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"

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
	//test input interfase
	if r == nil {
		return ASCII, nil
	}
	//make slice of byte from input reader
	buf, err := bufio.NewReader(r).Peek(ReadBufSize)
	if (err != nil) && (err != io.EOF) {
		return ASCII, err
	}

	//clear all counts and matching result
	//CodepageDic - global var and need cleaning befor reuse
	CodepageDic.clear()

	//match code page from BOM, support: utf-8, utf-16le, utf-16be, utf-32le or utf-32be
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
	return CodepageDic.Match(content)
}

//FileConvertCodePage - replace code page text file from one to another
// support convert only from/to Windows1251/IBM866
func FileConvertCodePage(fileName string, fromCP, toCP IDCodePage) error {
	if fromCP == toCP {
		return nil
	}

	if (fromCP != CP1251) && (fromCP != CP866) {
		return nil
	}

	if (toCP != CP1251) && (toCP != CP866) {
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

//IsSeparator - return true if input rune is SPACE or PUNCT
func IsSeparator(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r)
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
	case CP1251:
		s, _, err = transform.String(charmap.Windows1251.NewDecoder(), s)
	}
	switch toCP {
	case CP866:
		s, _, err = transform.String(charmap.CodePage866.NewEncoder(), s)
	case CP1251:
		s, _, err = transform.String(charmap.Windows1251.NewEncoder(), s)
	}
	return s, err
}

// CodePageAsString - return name of char set with id codepage
// if codepage not exist - return ""
func CodePageAsString(codepage IDCodePage) string {
	return CodepageDic[codepage].name
}

//DecodeUTF16 - decode slice of byte from UTF16 to UTF8
func DecodeUTF16(b []byte) string {
	if len(b)%2 != 0 {
		return string(b)
	}
	u16s := make([]uint16, 1)
	ret := &bytes.Buffer{}
	b8buf := make([]byte, 4)
	for i := 0; i < len(b); i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}
	return ret.String()
}
