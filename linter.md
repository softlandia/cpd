code_pages.go:45:2: `name` is unused (structcheck)
	name     string        //name of code page
	^
cpd_test.go:37:2: `st` is unused (structcheck)
	st string     //stop string
	^
cpd_test.go:127:9: Error return value of `os.Link` is not checked (errcheck)
	os.Link(fp.Join("test_files/866to1251.txt"), fp.Join("test_files/866to1251.tmp"))
	       ^
