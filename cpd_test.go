package cpd

import (
	"os"
	"testing"
)

type tCodePageAsString struct {
	id IDCodePage
	s  string
}

var dCodePageAsString = []tCodePageAsString{
	{0, ""},
	{3, "ASCII"},
	{IBM866, "IBM866"},
	{Windows1251, "Windows1251"},
	{60000, ""},
}

func TestCodePageAsString(t *testing.T) {
	for i, v := range dCodePageAsString {
		s := CodePageAsString(v.id)
		if s != v.s {
			t.Errorf("<CodePageAsString> on test: %d return: %s, expected: %s", i, s, v.s)
		}
	}
}

//TestCodePageDetect - тестирование метода CodePageDetect
// проверки на входные параметры:
// 1. nil		входящий поток явный nil, параметр останова отсутствует
// 2. nil, "~"	входящий поток явный nil, параметр останова присутствует
// 3. входящий поток не инициализированный объект, проверка на передачу пустого интерфейса
// проверка работы осуществляется через FileCodePageDetect()
func TestCodePageDetect(t *testing.T) {
	_, err := CodePageDetect(nil)
	if err == nil {
		t.Errorf("<CodePageDetect> on input nil return error == nil, expect error != nil\n")
	}
	_, err = CodePageDetect(nil, "~")
	if err == nil {
		t.Errorf("<CodePageDetect> on input nil return error == nil, expect error != nil\n")
	}

	var data *os.File
	res, err := CodePageDetect(data, "~")
	if err == nil {
		t.Errorf("<CodePageDetect> on input nil return error != nil, data: %+v, res: %d, code page: %s\n", data, res, CodePageAsString(res))
	}
}

func TestFileCodePageDetectSimple(t *testing.T) {
	res, err := FileCodePageDetect("test_files\\866to1251.txt")
	if err != nil {
		t.Errorf("<FileCodePageDetect()> on file '866to1251.txt' err expected: nil, got: %s\n", err)
	}
	if res != IBM866 {
		t.Errorf("<FileCodePageDetect()> on file '866to1251.txt' expected: %s, got: %s\n", IBM866, res)
	}
	res, err = FileCodePageDetect("test_files\\866&1251.txt")
	if err != nil {
		t.Errorf("<FileCodePageDetect()> on file '866&1251.txt' err expected: nil, got: %s\n", err)
	}
	if res != Windows1251 {
		t.Errorf("<FileCodePageDetect()> on file '866&1251.txt' expected: %s, got: %s\n", Windows1251, res)
	}
}

func TestFileCodePageDetectUtf8Bom(t *testing.T) {
	res, err := FileCodePageDetect("test_files\\utf8wbom.txt")
	if err != nil {
		t.Errorf("<FileCodePageDetect()> on file 'utf8wbom.txt' err expected: nil, got: %s\n", err)
	}
	if res != UTF8 {
		t.Errorf("<FileCodePageDetect()> on file 'utf8wbom.txt' expected: %s, got: %s\n", UTF8, res)
	}
}

type tFileCodePageDetectTest struct {
	fn string     //filename
	st string     //stop string
	e  error      //
	r  IDCodePage //expected result
}

var dFileCodePageDetect = []tFileCodePageDetectTest{
	{"test_files\\utf16BEwbom.txt", "", nil, UTF16BE},         //file contain utf16 big endian with bom rune at start
	{"test_files\\utf16be-woBOM.txt", "", nil, UTF16BE},       //file contain utf16 big endian with out bom rune at start
	{"test_files\\utf16le-wBOM.txt", "", nil, UTF16LE},        //file contain utf16 liitle endian with bom rune at start
	{"test_files\\utf16le-woBOM.txt", "", nil, UTF16LE},       //file contain utf16 liitle endian with out bom rune at start
	{"test_files\\utf8-woBOM.txt", "", nil, UTF8},             //file contain utf8 with out bom rune at start
	{"test_files\\866&1251.txt", "~X~", nil, Windows1251},     //befor ~X~ file contain 866, after 1251
	{"test_files\\866&1251.txt", "", nil, Windows1251},        //file contain more 1251 then 866
	{"test_files\\noCodePage.txt", "", nil, ASCII},            //file contain rune only ASCII
	{"test_files\\empty_file.txt", "", nil, ASCII},            //file exist but empty, no error, return ASCII
	{"test_files\\rune_encode_error.txt", "", nil, ASCII},     //file contain special rune -> encode error, but detect NO error
	{"test_files\\rune_error_1251.txt", "", nil, Windows1251}, //file contain 1251 and special rune -> encode error, but detect NO error
	{"test_files\\utf8wbom.txt", "", nil, UTF8},               //file contain utf8 with bom rune at start
	{"test_files\\utf16LEwbom.txt", "", nil, UTF16LE},         //file contain utf16 little endian with bom rune at start
}

//FileCodePageDetect
func TestFileCodePageDetect(t *testing.T) {
	var (
		err error
		res IDCodePage
	)
	for _, d := range dFileCodePageDetect {
		if len(d.st) == 0 {
			res, err = FileCodePageDetect(d.fn)
		} else {
			res, err = FileCodePageDetect(d.fn, d.st)
		}
		if err != d.e {
			t.Errorf("<FileCodePageDetect> on file '%s' expected error:  '%v', got: '%v', ", d.fn, d.e, err)
		}
		if res != d.r {
			t.Errorf("<FileCodePageDetect> on file '%s' expected result: %s, got: %s", d.fn, d.r, res)
		}
	}

	_, err = FileCodePageDetect("-.-") //file "-.-" not exist
	if err == nil {
		t.Errorf("<FileCodePageDetect> on file '-.-' must return error, but return nil")
	}

	_, err = FileCodePageDetect("") //file "" not exist
	if err == nil {
		t.Errorf("<FileCodePageDetect> on file '' must return error, but return nil")
	}

}

//FileConvertCodePage
func TestFileConvertCodePage(t *testing.T) {
	err := FileConvertCodePage("", 0, 1)
	if err == nil {
		t.Errorf("<FileConvertCodePage> on empty file name expected error, got: %v", err)
	}

	err = FileConvertCodePage("", 0, 0)
	if err != nil {
		t.Errorf("<FileConvertCodePage> on fromCp == toCp expected error==nil, got: %v", err)
	}

	err = FileConvertCodePage("test_files\\rune_encode_error.txt", IBM866, Windows1251)
	if err == nil {
		t.Errorf("<FileConvertCodePage> expected error, got: %v", err)
	}

	os.Link("test_files\\866to1251.txt", "test_files\\866to1251.tmp")
	err = FileConvertCodePage("test_files\\866to1251.tmp", IBM866, Windows1251)
	if err != nil {
		t.Errorf("<FileConvertCodePage> expect no err, got: %v", err)
	}
	os.Remove("test_files\\866to1251.tmp")
}

//ConvertCodePage
func TestStrConvertCodePage(t *testing.T) {
	_, err := StrConvertCodePage("1234", IBM866, Windows1251)
	if err != nil {
		t.Errorf("<StrConvertCodePage> on test 1 return unexpected err: %v", err)
	}
	_, err = StrConvertCodePage("1234", Windows1251, IBM866)
	if err != nil {
		t.Errorf("<StrConvertCodePage> on test 2 return unexpected err: %v", err)
	}
	_, err = StrConvertCodePage("", IBM866, Windows1251)
	if err != nil {
		t.Errorf("<StrConvertCodePage> with empty string must return ERROR, but retrurn: %v", err)
	}
	_, err = StrConvertCodePage("1234", IBM866, IBM866)
	if err != nil {
		t.Errorf("<StrConvertCodePage> with equal fromCP and toCp must return nil, but retrurn: %v", err)
	}
}
