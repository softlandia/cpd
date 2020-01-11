package cpd

import (
	"fmt"
	"os"
	fp "path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
	{fp.Join("test_files/866&1251.txt"), "", nil, CP1251},                    //file contain more 1251 then 866
	{fp.Join("test_files/empty_file.txt"), "", nil, UTF8},                    //file exist but empty, no error, return ASCII
	{fp.Join("test_files/IBM866.txt"), "", nil, CP866},                       //file contain IBM866
	{fp.Join("test_files/ISO8859-5.txt"), "", nil, ISOLatinCyrillic},         //file contain ISO8859-5
	{fp.Join("test_files/KOI8-r.txt"), "", nil, KOI8R},                       //file contain KOI8
	{fp.Join("test_files/KOI8-r2.txt"), "", nil, KOI8R},                      //file contain KOI8
	{fp.Join("test_files/noCodePage.txt"), "", nil, UTF8},                    //file contain rune only ASCII
	{fp.Join("test_files/rune_encode_error.txt"), "", nil, ISOLatinCyrillic}, //file contain special rune -> encode error, but detect NO error
	{fp.Join("test_files/rune_error_1251.txt"), "", nil, CP1251},             //file contain 1251 and special rune -> encode error, but detect NO error
	{fp.Join("test_files/utf8.txt"), "", nil, UTF8},                          //file contain utf8 with out bom rune at start
	{fp.Join("test_files/utf8-wBOM.txt"), "", nil, UTF8},                     //file contain utf8 with bom prefix
	{fp.Join("test_files/utf8-woBOM.txt"), "", nil, UTF8},                    //file contain utf8 with out bom rune at start
	{fp.Join("test_files/utf16be-wBOM.txt"), "", nil, UTF16BE},               //file contain utf16 big endian with bom
	{fp.Join("test_files/utf16be-woBOM.txt"), "", nil, UTF16BE},              //file contain utf16 big endian without bom
	{fp.Join("test_files/utf16le-wBOM.txt"), "", nil, UTF16LE},               //file contain utf16 little endian with bom
	{fp.Join("test_files/utf16le-woBOM.txt"), "", nil, UTF16LE},              //file contain utf16 little endian without bom
	{fp.Join("test_files/utf32be-wBOM.txt"), "", nil, UTF32BE},               //file contain utf32 big endian with bom
	{fp.Join("test_files/utf32be-woBOM.txt"), "", nil, UTF32BE},              //file contain utf32 big endian without bom
	{fp.Join("test_files/utf32le-wBOM.txt"), "", nil, UTF32LE},               //file contain utf32 little endian with bom
	{fp.Join("test_files/utf32le-woBOM.txt"), "", nil, UTF32LE},              //file contain utf32 little endian without bom
	{fp.Join("test_files/Win1251.txt"), "", nil, CP1251},                     //file contain Windows1251
	{fp.Join("test_files/win1251_upper.txt"), "", nil, CP1251},               //file contain Windows1251
}

//FileCodePageDetect
func TestFileCodePageDetect(t *testing.T) {
	var (
		err error
		res IDCodePage
	)
	for _, d := range dFileCodePageDetect {
		res, err = FileCodePageDetect(d.fn)
		assert.Equal(t, err, d.e, fmt.Sprintf("<FileCodePageDetect> on file '%s' expected error:  '%v', got: '%v', ", d.fn, d.e, err))
		assert.Equal(t, res, d.r, fmt.Sprintf("<FileCodePageDetect> on file '%s' expected result: %s, got: %s", d.fn, d.r, res))
	}

	_, err = FileCodePageDetect("-.-") //file "-.-" not exist
	assert.NotNil(t, err, "<FileCodePageDetect> on file '-.-' must return error, but return nil")

	_, err = FileCodePageDetect("") //file "" not exist
	assert.NotNil(t, err, "<FileCodePageDetect> on file '' must return error, but return nil")
}

//TestCodePageDetect - тестирование метода CodePageDetect
// проверки на входные параметры:
// 1. nil		входящий поток явный nil, параметр останова отсутствует
// 2. nil, "~"	входящий поток явный nil, параметр останова присутствует
// 3. входящий поток не инициализированный объект, проверка на передачу пустого интерфейса
// проверка самой работы осуществляется через FileCodePageDetect()
func TestCodePageDetect(t *testing.T) {
	tmp, err := CodePageDetect(nil)
	assert.Nil(t, err, fmt.Sprintf("<CodePageDetect> on input nil return error != nil\n"))
	assert.Equal(t, tmp, ASCII, fmt.Sprintf("<CodePageDetect> on input nil return code page != ASCII\n"))
	tmp, err = CodePageDetect(nil, "~")
	assert.Nil(t, err, fmt.Sprintf("<CodePageDetect> on input nil return error != nil\n"))
	assert.Equal(t, tmp, ASCII, fmt.Sprintf("<CodePageDetect> on input nil return code page != ASCII\n"))

	var data *os.File
	res, err := CodePageDetect(data)
	assert.NotNil(t, err, fmt.Sprintf("<CodePageDetect> on empty io.Reader return error != nil, data: %+v, err: %v\n", data, err))
	assert.Equal(t, res, ASCII, fmt.Sprintf("<CodePageDetect> on empty io.Reader = %+v return code page %s != ASCII\n", data, res))

	res, err = CodePageDetect(strings.NewReader(""))
	assert.Nil(t, err, fmt.Sprintf("<CodePageDetect> on input \"\" return error: %v, expected nil\n", err))
	assert.Equal(t, res, UTF8, fmt.Sprintf("<CodePageDetect> on input \"\" return %s, expected ASCII\n", res))
}

//FileConvertCodePage
func TestFileConvertCodePage(t *testing.T) {
	err := FileConvertCodePage("", CP866, CP1251)
	assert.NotNil(t, err, fmt.Sprintf("<FileConvertCodePage> on empty file name expected error, got: %v", err))

	err = FileConvertCodePage("", CP866, CP866)
	assert.Nil(t, err, fmt.Sprintf("<FileConvertCodePage> on fromCp == toCp expected error==nil, got: %v", err))

	err = FileConvertCodePage("123", UTF8, CP866)
	assert.Nil(t, err, fmt.Sprintf("<FileConvertCodePage> on fromCp or toCp not Windows1251 or IBM866 expected error == nil, got: %v", err))

	err = FileConvertCodePage("123", CP866, UTF16LE)
	assert.Nil(t, err, fmt.Sprintf("<FileConvertCodePage> on fromCp or toCp not Windows1251 or IBM866 expected error == nil, got: %v", err))

	err = FileConvertCodePage(fp.Join("test_files/rune_encode_error.txt"), CP866, CP1251)
	assert.NotNil(t, err, fmt.Sprintf("<FileConvertCodePage> expected error, got: %v", err))

	os.Link(fp.Join("test_files/866to1251.txt"), fp.Join("test_files/866to1251.tmp"))
	err = FileConvertCodePage(fp.Join("test_files/866to1251.tmp"), CP866, CP1251)
	assert.Nil(t, err, fmt.Sprintf("<FileConvertCodePage> expect no err, got: %v", err))
	os.Remove(fp.Join("test_files/866to1251.tmp"))
}

//ConvertCodePage
func TestStrConvertCodePage(t *testing.T) {
	b := []byte{0x91, 0xE2, 0xE0}
	s, err := StrConvertCodePage(string(b), CP866, CP1251)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> on test 1 return err: %v unexpected nil", err))
	b = []byte{0xD1, 0xF2, 0xF0}
	assert.Equal(t, s, string(b), fmt.Sprintf("<StrConvertCodePage> on test CP866->CP1251 return string: %s, expected: 'Стр'", s))

	b = []byte{0xFF, 0xC9, 0xB8}
	s, err = StrConvertCodePage(string(b), CP1251, CP866)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> on test CP1251->CP866 return unexpected err: %v", err))
	b = []byte{0xEF, 0x89, 0xF1}
	assert.Equal(t, s, string(b), fmt.Sprintf("<StrConvertCodePage> on test CP1251->CP866 return string: %s, expected: 'яЙё'", s))

	s, err = StrConvertCodePage("", CP866, CP1251)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> with empty input string must return ERROR nil, but retrurn: %v", err))
	assert.Equal(t, s, "", fmt.Sprintf("<StrConvertCodePage> with empty input string must return empty, but retrurn: %s", err))

	s, err = StrConvertCodePage("1234", CP866, CP866)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> with equal fromCP and toCp must return nil, but retrurn: %v", err))
	assert.Equal(t, s, "1234", fmt.Sprintf("<StrConvertCodePage> with equal fromCP and toCp must return input string, but retrurn: %s", s))
}
