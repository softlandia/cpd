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
	{CP866, "CP866"},
	{CP1251, "CP1251"},
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

type tFileCodePageDetectTest struct {
	fn string     //filename
	st string     //stop string
	e  error      //
	r  IDCodePage //expected result
}

var dFileCodePageDetect = []tFileCodePageDetectTest{
	{"test_files\\866&1251.txt", "", nil, CP1251},                    //file contain more 1251 then 866
	{"test_files\\empty_file.txt", "", nil, UTF8},                    //file exist but empty, no error, return ASCII
	{"test_files\\IBM866.txt", "", nil, CP866},                       //file contain IBM866
	{"test_files\\ISO8859-5.txt", "", nil, ISOLatinCyrillic},         //file contain ISO8859-5
	{"test_files\\KOI8-r.txt", "", nil, KOI8R},                       //file contain KOI8
	{"test_files\\KOI8-r2.txt", "", nil, KOI8R},                      //file contain KOI8
	{"test_files\\noCodePage.txt", "", nil, UTF8},                    //file contain rune only ASCII
	{"test_files\\rune_encode_error.txt", "", nil, ISOLatinCyrillic}, //file contain special rune -> encode error, but detect NO error
	{"test_files\\rune_error_1251.txt", "", nil, CP1251},             //file contain 1251 and special rune -> encode error, but detect NO error
	{"test_files\\utf8.txt", "", nil, UTF8},                          //file contain utf8 with out bom rune at start
	{"test_files\\utf8-wbom.txt", "", nil, UTF8},                     //file contain utf8 with bom prefix
	{"test_files\\utf8-woBOM.txt", "", nil, UTF8},                    //file contain utf8 with out bom rune at start
	{"test_files\\utf16be-wBOM.txt", "", nil, UTF16BE},               //file contain utf16 big endian with bom
	{"test_files\\utf16be-woBOM.txt", "", nil, UTF16BE},              //file contain utf16 big endian without bom
	{"test_files\\utf16le-wBOM.txt", "", nil, UTF16LE},               //file contain utf16 little endian with bom
	{"test_files\\utf16le-woBOM.txt", "", nil, UTF16LE},              //file contain utf16 little endian without bom
	{"test_files\\utf32be-wBOM.txt", "", nil, UTF32BE},               //file contain utf32 big endian with bom
	{"test_files\\utf32be-woBOM.txt", "", nil, UTF32BE},              //file contain utf32 big endian without bom
	{"test_files\\utf32le-wBOM.txt", "", nil, UTF32LE},               //file contain utf32 little endian with bom
	{"test_files\\utf32le-woBOM.txt", "", nil, UTF32LE},              //file contain utf32 little endian without bom
	{"test_files\\Win1251.txt", "", nil, CP1251},                     //file contain Windows1251
	{"test_files\\win1251_upper.txt", "", nil, CP1251},               //file contain Windows1251
}

//FileCodePageDetect
func TestFileCodePageDetect(t *testing.T) {
	var (
		err error
		res IDCodePage
	)
	for _, d := range dFileCodePageDetect {
		res, err = FileCodePageDetect(d.fn)
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

//TestCodePageDetect - тестирование метода CodePageDetect
// проверки на входные параметры:
// 1. nil		входящий поток явный nil, параметр останова отсутствует
// 2. nil, "~"	входящий поток явный nil, параметр останова присутствует
// 3. входящий поток не инициализированный объект, проверка на передачу пустого интерфейса
// проверка самой работы осуществляется через FileCodePageDetect()
func TestCodePageDetect(t *testing.T) {
	tmp, err := CodePageDetect(nil)
	if (err != nil) && (tmp != ASCII) {
		t.Errorf("<CodePageDetect> on input nil return error != nil or code page != ASCII\n")
	}
	tmp, err = CodePageDetect(nil, "~")
	if (err != nil) && (tmp != ASCII) {
		t.Errorf("<CodePageDetect> on input nil return error != nil or code page != ASCII\n")
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
	if res != CP866 {
		t.Errorf("<FileCodePageDetect()> on file '866to1251.txt' expected: %s, got: %s\n", CP866, res)
	}
	res, err = FileCodePageDetect("test_files\\866&1251.txt")
	if err != nil {
		t.Errorf("<FileCodePageDetect()> on file '866&1251.txt' err expected: nil, got: %s\n", err)
	}
	if res != CP1251 {
		t.Errorf("<FileCodePageDetect()> on file '866&1251.txt' expected: %s, got: %s\n", CP1251, res)
	}
}

func TestFileCodePageDetectUtf8Bom(t *testing.T) {
	res, err := FileCodePageDetect("test_files\\utf8-wBOM.txt")
	if err != nil {
		t.Errorf("<FileCodePageDetect()> on file 'utf8-wBOM.txt' err expected: nil, got: %s\n", err)
	}
	if res != UTF8 {
		t.Errorf("<FileCodePageDetect()> on file 'utf8-wBOM.txt' expected: %s, got: %s\n", UTF8, res)
	}
}

//FileConvertCodePage
func TestFileConvertCodePage(t *testing.T) {
	err := FileConvertCodePage("", CP866, CP1251)
	if err == nil {
		t.Errorf("<FileConvertCodePage> on empty file name expected error, got: %v", err)
	}

	err = FileConvertCodePage("", CP866, CP866)
	if err != nil {
		t.Errorf("<FileConvertCodePage> on fromCp == toCp expected error==nil, got: %v", err)
	}

	err = FileConvertCodePage("123", UTF8, CP866)
	if err != nil {
		t.Errorf("<FileConvertCodePage> on fromCp or toCp not Windows1251 or IBM866 expected error == nil, got: %v", err)
	}

	err = FileConvertCodePage("123", CP866, UTF16LE)
	if err != nil {
		t.Errorf("<FileConvertCodePage> on fromCp or toCp not Windows1251 or IBM866 expected error == nil, got: %v", err)
	}

	err = FileConvertCodePage("test_files\\rune_encode_error.txt", CP866, CP1251)
	if err == nil {
		t.Errorf("<FileConvertCodePage> expected error, got: %v", err)
	}

	os.Link("test_files\\866to1251.txt", "test_files\\866to1251.tmp")
	err = FileConvertCodePage("test_files\\866to1251.tmp", CP866, CP1251)
	if err != nil {
		t.Errorf("<FileConvertCodePage> expect no err, got: %v", err)
	}
	os.Remove("test_files\\866to1251.tmp")
}

//ConvertCodePage
func TestStrConvertCodePage(t *testing.T) {
	_, err := StrConvertCodePage("1234", CP866, CP1251)
	if err != nil {
		t.Errorf("<StrConvertCodePage> on test 1 return unexpected err: %v", err)
	}
	_, err = StrConvertCodePage("1234", CP1251, CP866)
	if err != nil {
		t.Errorf("<StrConvertCodePage> on test 2 return unexpected err: %v", err)
	}
	_, err = StrConvertCodePage("", CP866, CP1251)
	if err != nil {
		t.Errorf("<StrConvertCodePage> with empty string must return ERROR, but retrurn: %v", err)
	}
	_, err = StrConvertCodePage("1234", CP866, CP866)
	if err != nil {
		t.Errorf("<StrConvertCodePage> with equal fromCP and toCp must return nil, but retrurn: %v", err)
	}
}
