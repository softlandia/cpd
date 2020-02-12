package cpd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	fp "path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tStringHasBom struct {
	id IDCodePage
	s  string
	r  bool
}

var dStringHasBom = []tStringHasBom{
	{0, "", false},
	{ASCII, "", false},
	{CP866, "CP866", false},
	{CP1251, "\xD0\xEE\xF1\xF1\xE8\xFF", false},
	{CP1251, "\xff\xfe\xD0\xEE\xF1\xF1\xE8\xFF", false},   // contain UTF16LE bom, false because CP1251 have no bom
	{UTF8, "\xef\xbb\xbf\xD0\x94\xD0\xB5\xD0\xB4", true},  // Дед UTF8 with bom
	{UTF8, "\xef\xbb\xbe\xD0\xEE\xF1\xF1\xE8\xFF", false}, // UTF8 without bom
	{UTF8, "\xff\xbb\xbe\xD0\xEE\xF1\xF1\xE8\xFF", false}, // UTF8 without bom
	{UTF16BE, "\xfe\xff\xD0\xEE\xF1\xF1\xE8\xFF", true},
	{UTF16LE, "\xff\xfe\xD0\xEE\xF1\xF1\xE8\xFF", true},
	{UTF32BE, "\x00\x00\xfe\xff\xD0\xEE\xF1\xF1\xE8\xFF", true},
	{UTF32LE, "\xff\xfe\x00\x00\xD0\xEE\xF1\xF1\xE8\xFF", true},
}

func TestStringHasBom(t *testing.T) {
	for i, v := range dStringHasBom {
		assert.Equal(t, v.r, v.id.StringHasBom(v.s), fmt.Sprintf("test: %d, cp: %s", i, v.id))
	}
}

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
		s := CodepageAsString(v.id)
		assert.Equal(t, v.s, s, fmt.Sprintf("<CodePageAsString> on test: %d return: %s, expected: %s", i, s, v.s))
	}
}

func TestCodepageString(t *testing.T) {
	for i, v := range dCodePageAsString {
		s := v.id.String()
		assert.Equal(t, v.s, s, fmt.Sprintf("<CodePageAsString> on test: %d return: %s, expected: %s", i, s, v.s))
	}
}

type tFileCodePageDetectTest struct {
	fn string     //filename
	st string     //stop string, not using now
	e  error      //
	r  IDCodePage //expected result
}

var dFileCodePageDetect = []tFileCodePageDetectTest{
	{fp.Join("test_files/utf16le-woBOM-only-ru.txt"), "", nil, UTF16LE},      //file contain utf16 little endian without bom
	{fp.Join("test_files/utf16le-woBOM-no-ru.txt"), "", nil, UTF16LE},        //file contain utf16 little endian without bom
	{fp.Join("test_files/utf16le-woBOM-only-latin.txt"), "", nil, UTF16LE},   //file contain utf16 little endian without bom
	{fp.Join("test_files/utf16le_las.txt"), "", nil, UTF16LE},                //file contain utf16 little endian without bom
	{fp.Join("test_files/utf16be_las.txt"), "", nil, UTF16BE},                //file contain utf16 big endian with bom
	{fp.Join("test_files/866&1251.txt"), "", nil, CP1251},                    //file contain more 1251 then 866
	{fp.Join("test_files/empty_file.txt"), "", nil, UTF8},                    //file exist but empty, no error, return ASCII
	{fp.Join("test_files/IBM866.txt"), "", nil, CP866},                       //file contain IBM866
	{fp.Join("test_files/ISO8859-5.txt"), "", nil, ISOLatinCyrillic},         //file contain ISO8859-5
	{fp.Join("test_files/KOI8-r.txt"), "", nil, KOI8R},                       //file contain KOI8
	{fp.Join("test_files/KOI8-r2.txt"), "", nil, KOI8R},                      //file contain KOI8
	{fp.Join("test_files/noCodePage.txt"), "", nil, UTF8},                    //file contain rune only ASCII
	{fp.Join("test_files/rune_encode_error.txt"), "", nil, ISOLatinCyrillic}, //file contain special rune -> encode error, but detect NO error
	{fp.Join("test_files/rune_error_1251.txt"), "", nil, CP1251},             //file contain 1251 and special rune -> encode error, but detect NO error
	{fp.Join("test_files/utf8.txt"), "", nil, UTF8},                          //file contain utf8 without bom
	{fp.Join("test_files/utf8-wBOM.txt"), "", nil, UTF8},                     //file contain utf8 with bom
	{fp.Join("test_files/utf8-woBOM.txt"), "", nil, UTF8},                    //file contain utf8 without bom
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
	{fp.Join("test_files/utf16be-woBOM-only-latin.txt"), "", nil, UTF16BE},   //file contain utf16be with bom
	{fp.Join("test_files/utf16be-woBOM-no-ru.txt"), "", nil, UTF16BE},        //file contain utf16be with bom
	{fp.Join("test_files/utf16be-woBOM-only-ru.txt"), "", nil, UTF16BE},      //file contain utf16be with bom
	{fp.Join("test_files/utf32be-ascii-no-ru.txt"), "", nil, UTF16BE},        //file contain utf32be w/o bom w/o ru, detected as UTF16BE
	{fp.Join("test_files/utf32le-ascii-no-ru.txt"), "", nil, UTF16LE},        //file contain utf32le w/o bom w/o ru, detected as UTF16LE
}

//FileCodePageDetect
func TestFileCodePageDetect(t *testing.T) {
	var (
		err error
		res IDCodePage
	)
	for _, d := range dFileCodePageDetect {
		res, err = FileCodepageDetect(d.fn)
		assert.Equal(t, err, d.e, fmt.Sprintf("<FileCodePageDetect> on file '%s' expected error:  '%v', got: '%v', ", d.fn, d.e, err))
		assert.Equal(t, res, d.r, fmt.Sprintf("<FileCodePageDetect> on file '%s' expected result: %s, got: %s", d.fn, d.r, res))
	}

	_, err = FileCodepageDetect("-.-") //file "-.-" not exist
	assert.NotNil(t, err, "<FileCodePageDetect> on file '-.-' must return error, but return nil")

	_, err = FileCodepageDetect("") //file "" not exist
	assert.NotNil(t, err, "<FileCodePageDetect> on file '' must return error, but return nil")
}

func fileCodepageDetect(wg *sync.WaitGroup, t *testing.T, trusted IDCodePage, f string) {
	defer wg.Done()
	res, _ := FileCodepageDetect(f)
	assert.Equal(t, trusted, res, fmt.Sprintf("<FileCodePageDetect> on file '%s' expected: %s, got: %s", f, trusted, res))
}

//тестирование в многопоточном режиме
func TestFileCodePageDetectM(t *testing.T) {
	var wg sync.WaitGroup
	for _, d := range dFileCodePageDetect {
		wg.Add(1)
		go fileCodepageDetect(&wg, t, d.r, d.fn)
	}
	wg.Wait()
}

// TestCodePageDetect - тестирование метода CodePageDetect
// проверки на входные параметры
// проверка самой работы осуществляется через FileCodePageDetect()
func TestCodePageDetect(t *testing.T) {
	// 1. nil		входящий поток явный nil, параметр останова отсутствует
	tmp, err := CodepageDetect(nil)
	assert.Nil(t, err, fmt.Sprintf("<CodePageDetect> on input nil return error != nil\n"))
	assert.Equal(t, tmp, ASCII, fmt.Sprintf("<CodePageDetect> on input nil return code page != ASCII\n"))

	var data *os.File
	// 2. nil, "~"	входящий поток явный nil, параметр останова присутствует
	res, err := CodepageDetect(data)
	assert.NotNil(t, err, fmt.Sprintf("<CodePageDetect> on empty io.Reader return error != nil, data: %+v, err: %v\n", data, err))
	assert.Equal(t, res, ASCII, fmt.Sprintf("<CodePageDetect> on empty io.Reader = %+v return code page %s != ASCII\n", data, res))

	res, err = CodepageDetect(strings.NewReader(""))
	assert.Nil(t, err, fmt.Sprintf("<CodePageDetect> on input \"\" return error: %v, expected nil\n", err))
	assert.Equal(t, res, UTF8, fmt.Sprintf("<CodePageDetect> on input \"\" return %s, expected ASCII\n", res))
}

//FileConvertCodePage
func TestFileConvertCodePage(t *testing.T) {
	err := FileConvertCodepage("", CP866, CP1251)
	assert.NotNil(t, err, fmt.Sprintf("<FileConvertCodePage> on empty file name expected error, got: %v", err))

	err = FileConvertCodepage("", CP866, CP866)
	assert.Nil(t, err, fmt.Sprintf("<FileConvertCodePage> on fromCp == toCp expected error==nil, got: %v", err))

	err = FileConvertCodepage("123", UTF8, CP866)
	assert.Nil(t, err, fmt.Sprintf("<FileConvertCodePage> on fromCp or toCp not Windows1251 or IBM866 expected error == nil, got: %v", err))

	err = FileConvertCodepage("123", CP866, UTF16LE)
	assert.Nil(t, err, fmt.Sprintf("<FileConvertCodePage> on fromCp or toCp not Windows1251 or IBM866 expected error == nil, got: %v", err))

	err = FileConvertCodepage(fp.Join("test_files/rune_encode_error.txt"), CP866, CP1251)
	assert.NotNil(t, err, fmt.Sprintf("<FileConvertCodePage> expected error, got: %v", err))

	os.Link(fp.Join("test_files/866to1251.txt"), fp.Join("test_files/866to1251.tmp"))
	err = FileConvertCodepage(fp.Join("test_files/866to1251.tmp"), CP866, CP1251)
	assert.Nil(t, err, fmt.Sprintf("<FileConvertCodePage> expect no err, got: %v", err))
	os.Remove(fp.Join("test_files/866to1251.tmp"))
}

//ConvertCodePage
func TestStrConvertCodePage(t *testing.T) {
	s, err := StrConvertCodepage(string([]byte{0x91, 0xE2, 0xE0}), CP866, CP1251)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> on test 1 return err: %v unexpected nil", err))
	assert.Equal(t, s, string([]byte{0xD1, 0xF2, 0xF0}), fmt.Sprintf("<StrConvertCodePage> on test CP866->CP1251 return string: %s, expected: 'Стр'", s))

	s, err = StrConvertCodepage(string([]byte{0xFF, 0xC9, 0xB8}), CP1251, CP866)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> on test CP1251->CP866 return unexpected err: %v", err))
	assert.Equal(t, s, string([]byte{0xEF, 0x89, 0xF1}), fmt.Sprintf("<StrConvertCodePage> on test CP1251->CP866 return string: %s, expected: 'яЙё'", s))

	s, err = StrConvertCodepage("", CP866, CP1251)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> with empty input string must return ERROR nil, but return: %v", err))
	assert.Equal(t, s, "", fmt.Sprintf("<StrConvertCodePage> with empty input string must return empty, but return: %s", err))

	s, err = StrConvertCodepage("1234", CP866, CP866)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> with equal fromCP and toCp must return nil, but retrurn: %v", err))
	assert.Equal(t, s, "1234", fmt.Sprintf("<StrConvertCodePage> with equal fromCP and toCp must return input string, but return: %s", s))

	s, err = StrConvertCodepage(string([]byte{0xD0, 0xEE, 0xF1, 0xF1, 0xE8, 0xFF}), CP1251, UTF8)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> from CP1251 to UTF expect nil, got: %v", err))
	assert.Equal(t, s, "Россия", fmt.Sprintf("<StrConvertCodePage> '%s' wrong return from CP1251 to UTF string", s))

	s, err = StrConvertCodepage(string([]byte{0x90, 0xAE, 0xE1, 0xE1, 0xA8, 0xEF}), CP866, UTF8)
	assert.Nil(t, err, fmt.Sprintf("<StrConvertCodePage> from CP866 to UTF expect nil, got: %v", err))
	assert.Equal(t, s, "Россия", fmt.Sprintf("<StrConvertCodePage> '%s' wrong return from CP866 to UTF string", s))
}

type tDecodeUTF16le struct {
	iStr string
	oStr string
}

var dDecodeUTF16le = []tDecodeUTF16le{
	{"", ""}, // empty string
	{"\xFF\xFE\x14\x04\x35\x04\x34\x04", "Дед"}, // UTF16le with BOM	1
	{"\x14\x04\x35\x04\x34\x04", "Дед"},         // UTF16le w/o BOM		2
}

func TestDecodeUtf16le(t *testing.T) {
	for i, d := range dDecodeUTF16le {
		s := DecodeUTF16le(d.iStr)
		assert.Equal(t, d.oStr, s, fmt.Sprintf("test #%d not pass", i))
	}
}

type tDecodeUTF16be struct {
	iStr string
	oStr string
}

var dDecodeUTF16be = []tDecodeUTF16le{
	{"", ""}, // empty string
	{"\xFE\xFF\x04\x14\x04\x35\x04\x34", "Дед"}, // UTF16be with BOM	1
	{"\x04\x14\x04\x35\x04\x34", "Дед"},         // UTF16be w/o BOM		2
}

func TestDecodeUtf16be(t *testing.T) {
	for i, d := range dDecodeUTF16be {
		s := DecodeUTF16be(d.iStr)
		assert.Equal(t, d.oStr, s, fmt.Sprintf("test #%d not pass", i))
	}
}

type tNewReader struct {
	iStr   string // input string
	oStr   string // result string
	cpName string // codepage name
}

//здесь все тесты без ошибок
//
var dNewReader = []tNewReader{
	{"\xef\xbb\xbf\xD0\x94\xD0\xB5\xD0\xB4", "Дед", ""},                // UTF8 with BOM     0
	{"\xD0\x94\xD0\xB5\xD0\xB4", "Дед", ""},                            // UTF8 w/o BOM      1
	{"\xFE\xFF\x04\x14\x04\x35\x04\x34", "Дед", ""},                    // UTF16be with BOM  2
	{"\x04\x14\x04\x35\x04\x34\x00\x20\x04\x20\x04\x43", "Дед Ру", ""}, // UTF16be w/o BOM   3
	{"\x04\x14\x04\x35\x04\x34", "Дед", "UTF-16BE"},                    // UTF16be w/o BOM   4
	{"\xFF\xFE\x14\x04\x35\x04\x34\x04", "Дед", ""},                    // UTF16le with BOM  5
	{"\x14\x04\x35\x04\x34\x04\x20\x00\x20\x04\x43\x04", "Дед Ру", ""}, // UTF16le w/o BOM   6
	{"\x14\x04\x35\x04\x34\x04", "Дед", "UTF16le"},                     // UTF16le w/o BOM   7
	{"\xD0\xEE\xF1\xF1\xE8\xFF", "Россия", ""},                         // CP1251            8
	{"\xCE\xED\xEE", "Оно", ""},                                        // CP1251            9
	{"\xCE\xED\xEE", "Оно", "w1251"},                                   // CP1251            10
	{"\xCE\xED\xEE", "Оно", "win1251"},                                 // CP1251            11
	{"1", "1", ""},                                                     // ascii string      12
	{"1", "1", "CP866"},                                                // ascii string      13
	{"1", "1", "CP/866"},                                               // ascii string      14
	{"", "", "CP1251"},                                                 // ascii string      15
	{"", "", "CP"},                                                     // ascii string      16
	{"", "", ""},                                                       // ascii string      17
}

func TestNewReader(t *testing.T) {
	var (
		r   io.Reader
		err error
	)
	for i, d := range dNewReader {
		tmp := strings.NewReader(d.iStr)
		if len(d.cpName) > 0 {
			r, err = NewReader(tmp, d.cpName)
		} else {
			r, err = NewReader(tmp)
		}
		assert.Nil(t, err, fmt.Sprintf("test #%d not pass, NewReader() return error", i))
		b, err := ioutil.ReadAll(r)
		assert.Nil(t, err, fmt.Sprintf("test #%d not pass, ReadAll() return error", i))
		assert.Equal(t, d.oStr, string(b), fmt.Sprintf("test #%d not pass", i))
	}
	//дополнительно тестирование nil ридера
	r, err = NewReader(nil)
	assert.Equal(t, err.Error(), "cpd: input reader is nil", "test NewReader(nil) return err==nil, expect: 'cpd: input reader is nil'")
	assert.Nil(t, r, "test NewReader(nil) return r!=nil, expect: nil")
}

var dNewReaderErrors = []tNewReader{
	{"\x00\x00\x04\x20\x00\x00\x04\x43\x00\x00\x04\x41\x00\x00\x04\x41", "cpd: codepage not support decode", "UTF-32Be"}, // UTF-32be Русс	  0
	{"\x00\x00\x04\x20\x00\x00\x04\x43\x00\x00\x04\x41\x00\x00\x04\x41", "cpd: codepage not support decode", ""},         // UTF-32be Русс	  1
	{"\x04\x20\x00\x00\x04\x43\x00\x00\x04\x41\x00\x00\x04\x41\x00\x00", "cpd: codepage not support decode", "UTF32le"},  // UTF-32le Русс	  2
	{"\x04\x20\x00\x00\x04\x43\x00\x00\x04\x41\x00\x00\x04\x41\x00\x00", "cpd: codepage not support decode", ""},         // UTF-32le Русс	  3
}

// тестирование случаев с ошибкой
func TestNewReaderError(t *testing.T) {
	var (
		r   io.Reader
		err error
	)
	for i, d := range dNewReaderErrors {
		tmp := strings.NewReader(d.iStr)
		if len(d.cpName) > 0 {
			r, err = NewReader(tmp, d.cpName)
		} else {
			r, err = NewReader(tmp)
		}
		// мы должны получить ошибку: 'cpd: codepage not support decode'
		// кроме того вернётся поданый на вход ридер
		assert.NotNil(t, err, fmt.Sprintf("test #%d not pass, NewReader() return nil error", i))
		assert.Equal(t, tmp, r, fmt.Sprintf("test #%d not pass", i))
	}
}

type tNewReaderTo struct {
	iStr   string // input string
	oStr   string // result string
	cpName string // codepage name
}

var dNewReaderTo = []tNewReaderTo{
	{"Дед", "\x84\xA5\xA4", "CP866"},                                                   // 0
	{"Дед", "\x04\x14\x04\x35\x04\x34", "UTF16be"},                                     // 1
	{"Дед Ру", "\x04\x14\x04\x35\x04\x34\x00\x20\x04\x20\x04\x43", "UTF-16be"},         // 2
	{"Дед Ру", "\x14\x04\x35\x04\x34\x04\x20\x00\x20\x04\x43\x04", "UTF16Le"},          // 3
	{"Дед", "\x14\x04\x35\x04\x34\x04", "UTF16le"},                                     // 4
	{"Оно", "\xCE\xED\xEE", "CP-1251"},                                                 // 5
	{"О н о", "\xCE\x20\xED\x20\xEE", "win1251"},                                       // 6
	{"\xef\xbb\xbf\xD0\x9E\x20\xD0\xBD\x20\xD0\xBe", "\xCE\x20\xED\x20\xEE", "CP1251"}, // 7 input UTF8 with bom, convert to CP1251, BOM is expected to be deleted
	{"\xef\xbb\xbf\xD0\x94\xD0\xB5\xD0\xB4", "Дед", "UTF8"},                            // 8 input UTF8 with bom, convert to UTF8, BOM is expected to be deleted
	{"", "", "CP866"}, // empty input string but not empty exist codepage name  11
}

func TestNewReaderTo(t *testing.T) {
	for i, d := range dNewReaderTo {
		r, err := NewReaderTo(strings.NewReader(d.iStr), d.cpName)
		assert.Nil(t, err, fmt.Sprintf("test #%d not pass, NewReaderTo() return error: '%v'", i, err))
		b, err := ioutil.ReadAll(r)
		assert.Nil(t, err, fmt.Sprintf("test #%d not pass, ReadAll() return error: '%v'", i, err))
		assert.Equal(t, d.oStr, string(b), fmt.Sprintf("test #%d not pass", i))
	}
}

// тестирование случаев с ошибкой
var dNewReaderToError = []tNewReaderTo{
	{"sample", "cpd: output codepage not support encode", "UTF-32"}, // 0 landing codepage not support encoding
	{"1234", "cpd: output codepage not support encode", ""},         // 2 empty landing codepage submitted
	{"Дед Ру", "cpd: output codepage not support encode", "cpMY"},   // 3 not exist landing codepage submitted
	{"", "cpd: output codepage not support encode", "<CP>"},         // 4 empty input string but not empty landing codepage submitted
	{"", "cpd: output codepage not support encode", ""},             // 5 empty input string and empty landing codepage submitted
}

func TestNewReaderToError(t *testing.T) {
	for i, d := range dNewReaderToError {
		_, err := NewReaderTo(strings.NewReader(d.iStr), d.cpName)
		assert.Equal(t, d.oStr, err.Error(), fmt.Sprintf("test #%d. error expect: %s, got: %s", i, d.oStr, err.Error()))
	}
}

type tSupportedEncoders struct {
	cpn string // input codepage name
	res bool   // result
}

var dSupportedEncoders = []tSupportedEncoders{
	{"", false},        // 0
	{"<>", false},      // 1
	{"CP866", true},    // 2
	{"UTF32", false},   // 3
	{"UTF32le", false}, // 4
}

func TestSupportedEncoder(t *testing.T) {
	for i, d := range dSupportedEncoders {
		r := SupportedEncoder(d.cpn)
		assert.Equal(t, d.res, r, fmt.Sprintf("test #%d not pass", i))
	}
}
