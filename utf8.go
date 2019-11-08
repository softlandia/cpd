package cpd

import "encoding/binary"

//unit for UTF8

func runesMatchUTF8(data []byte, tbl *codePageTable) (counts int) {
	n := len(data)/2 - 1
	if n <= 0 {
		return
	}
	for i := 0; i < n; i += 2 {
		t := data[i : i+2]
		d := binary.BigEndian.Uint16(t)
		j := tbl.containsRune(rune(d))
		if j > 0 {
			(*tbl)[j].count++
			//counts++
		}
		if isUTF8(rune(d)) {
			counts++
		}
	}
	return
}

const (
	cpUTF8StartUpperChar = 0xD090
	cpUTF8StopUpperChar  = 0xD0AF
	cpUTF8StartLowerChar = 0xD0B0
	cpUTF8StopLowerChar  = 0xD18F
)

func isUpperUTF8(r rune) bool {
	return (r >= cpUTF8StartUpperChar) && (r <= cpUTF8StopUpperChar)
}

func isLowerUTF8(r rune) bool {
	return (r >= cpUTF8StartLowerChar) && (r <= cpUTF8StopLowerChar)
}

func isUTF8(r rune) bool {
	return isUpperUTF8(r) || isLowerUTF8(r)
}
