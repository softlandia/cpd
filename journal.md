v0.3.3		2020.01.17	add to type CodePage field BOM, 
				refactoring IDCodePage.String() - get name from global var CodepageDic
				add IDCodePage.StringHasBom(s string) bool - checks that the string s contains characters BOM matching this given codepage
				add IDCodePage.DeleteBom(s string) string - return string without prefix bom bytes
#TODO
   1 UTF16LE & UTF16BE not recognized correctly if file no contains russian characters
