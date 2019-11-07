package cpd

//checkHeader - check buffer for match to utf-8, utf-16le or utf-16be BOM
func checkHeader(b []byte) (id IDCodePage, res bool) {
	if bomUTF8(b) {
		return UTF8, true
	}
	if bomUTF16le(b) {
		return UTF16LE, true
	}
	if bomUTF16be(b) {
		return UTF16BE, true
	}
	return ASCII, false
}

func bomUTF8(b []byte) bool {
	if len(b) < 3 {
		return false
	}
	return (b[0] == 0xEF) && (b[1] == 0xBB) && (b[2] == 0xBF)
}

func bomUTF16le(b []byte) bool {
	if len(b) < 2 {
		return false
	}
	return (b[0] == 0xFF) && (b[1] == 0xFE)
}

func bomUTF16be(b []byte) bool {
	if len(b) < 2 {
		return false
	}
	return (b[0] == 0xFE) && (b[1] == 0xFF)
}

//ASCII block
func itASCII(r rune, tbl *codePageTable) int {
	return 0
}

func runesMatchASCII(b []byte, tbl *codePageTable) int {
	return 0
}
