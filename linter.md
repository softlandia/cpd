cpTable.go:33:19: func `(*cpTable).sort` is unused (unused)
utf8.go:113:6: func `toUTF8` is unused (unused)
code_pages.go:171:24: func `TCodepagesDic.clear` is unused (unused)
cpTable.go:27:19: func `(*cpTable).clear` is unused (unused)
sample\main.go:16:14: Error return value of `FindFilesExt` is not checked (errcheck)
	FindFilesExt(&fl, ".\\", os.Args[1])
	            ^
cpd_test.go:183:9: Error return value of `os.Link` is not checked (errcheck)
	os.Link(fp.Join("test_files/866to1251.txt"), fp.Join("test_files/866to1251.tmp"))
	       ^
cpd_test.go:70:2: `st` is unused (structcheck)
	st string     //stop string, not using now
	^
