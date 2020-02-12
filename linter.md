utf8.go:113:6: `toUTF8` is unused (deadcode)
func toUTF8(s string) string {
     ^
cpd_test.go:235:6: `tDecodeUTF16be` is unused (deadcode)
type tDecodeUTF16be struct {
     ^
code_pages.go:60:9: Error return value of `r.Read` is not checked (errcheck)
		r.Read(make([]byte, UTF8.BomLen())) // считываем в никуда количество байт занимаемых BOM этой кодировки
		      ^
cpd.go:183:17: Error return value of `tmpReader.Read` is not checked (errcheck)
		tmpReader.Read(make([]byte, cp.BomLen())) // считываем в никуда количество байт занимаемых BOM этой кодировки
		              ^
cpd_test.go:184:9: Error return value of `os.Link` is not checked (errcheck)
	os.Link(fp.Join("test_files/866to1251.txt"), fp.Join("test_files/866to1251.tmp"))
	       ^
sample\detect-all-files\main.go:18:14: Error return value of `FindFilesExt` is not checked (errcheck)
	FindFilesExt(&fl, ".\\", os.Args[1])
	            ^
cpd_test.go:236:2: `iStr` is unused (structcheck)
	iStr string
	^
cpd_test.go:237:2: `oStr` is unused (structcheck)
	oStr string
	^
cpd_test.go:72:2: `st` is unused (structcheck)
	st string     //stop string, not using now
	^
cpTable.go:33:19: func `(*cpTable).sort` is unused (unused)
code_pages.go:210:24: func `TCodepagesDic.clear` is unused (unused)
cpTable.go:27:19: func `(*cpTable).clear` is unused (unused)
