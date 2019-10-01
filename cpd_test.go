package cpd

import (
	"os"
	"testing"

	"github.com/softlandia/cpd/internal/cp"
)

type tCodePageAsString struct {
	id uint16
	s  string
}

var dCodePageAsString = []tCodePageAsString{
	{0, ""},
	{3, "ASCII"},
	{cp.IBM866, "IBM866"},
	{cp.Windows1251, "Windows1251"},
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

//CodePageDetect
func TestCodePageDetect(t *testing.T) {
	res, err := CodePageDetect("test_files\\866&1251.txt", "~X~") //befor ~X~ file contain 866, after 1251
	if err != nil {
		t.Errorf("<CodePageDetect> on file '%s' return error: %v", "866&1251.txt", err)
	}
	if res != cp.IBM866 {
		t.Errorf("<CodePageDetect> on file '%s' expected 866 got: %s", "866&1251.txt", CodePageAsString(res))
	}

	res, err = CodePageDetect("test_files\\866&1251.txt") //file contain more 1251 then 866
	if res != cp.Windows1251 {
		t.Errorf("<CodePageDetect> on file '%s' expected 1251 got: %s", "866&1251.txt", CodePageAsString(res))
	}

	_, err = CodePageDetect("-.-") //file "-.-" not exist
	if err == nil {
		t.Errorf("<CodePageDetect> on file '-.-' must return error, but return nil")
	}

	res, _ = CodePageDetect("test_files\\noCodePage.txt") //file contain rune only ASCII
	if res != cp.ASCII {
		t.Errorf("<CodePageDetect> on file 'noCodePage.txt' expect ASCII got: %s", CodePageAsString(res))
	}

	res, err = CodePageDetect("test_files\\empty_file.txt")
	if (res != cp.ASCII) || (err != nil) {
		t.Errorf("<CodePageDetect> on file 'empty_file.txt' expect ASCII and no error got: %s and %v", CodePageAsString(res), err)
	}

	res, err = CodePageDetect("test_files\\rune_encode_error.txt")
	if (res != cp.ASCII) || (err != nil) {
		t.Errorf("<CodePageDetect> on file 'rune_encode_error.txt' expect ASCII and no error got: %s and %v", CodePageAsString(res), err)
	}

	res, err = CodePageDetect("test_files\\rune_error_1251.txt")
	if res != cp.Windows1251 {
		t.Errorf("<CodePageDetect> on file 'rune_error_1251.txt' expect 1251 and no error got: %s and %v", CodePageAsString(res), err)
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

	err = FileConvertCodePage("test_files\\rune_encode_error.txt", cp.IBM866, cp.Windows1251)
	if err == nil {
		t.Errorf("<FileConvertCodePage> expected error, got: %v", err)
	}

	os.Link("test_files\\866to1251.txt", "test_files\\866to1251.tmp")
	err = FileConvertCodePage("test_files\\866to1251.tmp", cp.IBM866, cp.Windows1251)
	if err != nil {
		t.Errorf("<FileConvertCodePage> expect no err, got: %v", err)
	}
	os.Remove("test_files\\866to1251.tmp")
}

//ConvertCodePage
func TestStrConvertCodePage(t *testing.T) {
	_, err := StrConvertCodePage("1234", cp.IBM866, cp.Windows1251)
	if err != nil {
		t.Errorf("<StrConvertCodePage> on test 1 return unexpected err: %v", err)
	}
	_, err = StrConvertCodePage("1234", cp.Windows1251, cp.IBM866)
	if err != nil {
		t.Errorf("<StrConvertCodePage> on test 2 return unexpected err: %v", err)
	}
	_, err = StrConvertCodePage("", cp.IBM866, cp.Windows1251)
	if err != nil {
		t.Errorf("<StrConvertCodePage> with empty string must return ERROR, but retrurn: %v", err)
	}
	_, err = StrConvertCodePage("1234", cp.IBM866, cp.IBM866)
	if err != nil {
		t.Errorf("<StrConvertCodePage> with equal fromCP and toCp must return nil, but retrurn: %v", err)
	}
}
