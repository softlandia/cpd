//Package cpd - code page detect
// (c) 2019 softlandia@gmail.com
// v0.1.0
// 01/oct/2019
package cpd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/softlandia/cpd/internal/cp"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

//StrConvertCodePage - convert string from one code page to another
func StrConvertCodePage(s string, fromCP, toCP uint16) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if fromCP == toCP {
		return s, nil
	}

	var err error

	switch fromCP {
	case cp.IBM866:
		s, _, err = transform.String(charmap.CodePage866.NewDecoder(), s)
	case cp.Windows1251:
		s, _, err = transform.String(charmap.Windows1251.NewDecoder(), s)
	}
	switch toCP {
	case cp.IBM866:
		s, _, err = transform.String(charmap.CodePage866.NewEncoder(), s)
	case cp.Windows1251:
		s, _, err = transform.String(charmap.Windows1251.NewEncoder(), s)
	}
	return s, err
}

// CodePageAsString - return name of char set with id codepage
// if codepage not exist - return ""
func CodePageAsString(codepage uint16) string {
	return cp.Name[codepage]
}

//CodePageDetect - detect code page of file
//return 0 if code page can not be detected
//return const cpd.CpWindows1251 for Windows code page 1251
//return const cdp.CpIBM866 for IBM 866 code page
//return conts cdp.CpASCII by default or on error
//EF-BB-BF utf8 bom
func CodePageDetect(fn string, stopStr ...string) (uint16, error) {
	var (
		count1251 int //счётчик символов в кодировке 1251
		count866  int //счётчик символов в кодировке 866
	)

	iFile, err := os.Open(fn)
	if err != nil {
		return CpASCII, err
	}
	defer iFile.Close()

	iScanner := bufio.NewScanner(iFile)
	for i := 0; iScanner.Scan(); i++ {
		s := iScanner.Text()
		if (len(stopStr) > 0) && strings.Contains(s, stopStr[0]) { //stopStr[0] - строка при обнаружении которой останавливаемся, stopStr - слайс строк
			break
		}
		for j := range s {
			if isRune1251(rune(s[j])) { //проверка принадлежности символа позициям алфавитных символов в кодовой таблице 1251
				count1251++
			}
			if isRune866(rune(s[j])) { //проверка принадлежности символа позициям алфавитных символов в кодовой таблице 866
				count866++
			}
		}
	}
	switch {
	case count1251 > count866:
		return CpWindows1251, nil
	case count1251 < count866:
		return CpIBM866, nil
	}
	return CpASCII, nil
}

//FileConvertCodePage - replace code page text file from one to another
func FileConvertCodePage(fileName string, fromCP, toCP uint16) error {
	if fromCP == toCP {
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

const (
	cp866r1Min  = 0x80 //заглавная буква А
	cp866r1Max  = 0xAF //строчная буква п - в этом интервале в 866 раскладке лежит большинство русских букв
	cp866r2Min  = 0xE0 //строчная р
	cp866r2Max  = 0xF1 //строчна ё - в этом интервале лежат остальные русские буквы
	cp1251s1    = 0xA8 //Ё
	cp1251s2    = 0xB8 //ё в этой позиции в 866 лежит псевдографика
	cp1251r1Min = 0xC0 //с этой позиции начинается весь алфавит
	cp1251r1Max = 0xFF //заканчивается
	cpKOI8RMin  = 0xC0 //начало интервала
	cpKOI8RMax  = 0xFF //конец интервала
)

func isRune1251(r rune) bool {
	switch {
	case r == cp1251s1:
		return true
	case r == cp1251s2:
		return true
	case (r >= cp1251r1Min) && (r <= cp1251r1Max):
		return true
	}
	return false
}

func isRune866(r rune) bool {
	switch {
	case (r >= cp866r1Min) && (r <= cp866r1Max):
		return true
	case (r >= cp866r2Min) && (r <= cp866r2Max):
		return true
	}
	return false
}
