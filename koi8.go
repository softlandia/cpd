package cpd

import "unicode"

//unit for koi-8

func runesMatchKOI8(data []byte, tbl *codePageTable) (counts int) {
	for i := range data {
		if i < 2 {
			continue
		}
		//case " Us" - separator_UPPER_symbol
		if unicode.IsPunct(rune(data[i-2])) && isUpperKOI8(rune(data[i-1])) {
			j := tbl.containsRune(rune(data[i]))
			if j > 0 {
				(*tbl)[j].count++
				counts++
				continue
			}
		}
		if isKOI8(rune(data[i-1])) {
			j := tbl.containsRune(rune(data[i]))
			if j > 0 {
				(*tbl)[j].count++
				counts++
			}
		}
	}
	return
}

const (
	cpKOI8StartUpperChar = 0xE0
	cpKOI8StopUpperChar  = 0xFF
	cpKOI8StartLowerChar = 0xC0
	cpKOI8StopLowerChar  = 0xDF
)

func isUpperKOI8(r rune) bool {
	return (r >= cpKOI8StartUpperChar) && (r <= cpKOI8StopUpperChar)
}

func isLowerKOI8(r rune) bool {
	return (r >= cpKOI8StartLowerChar) && (r <= cpKOI8StopLowerChar)
}

func isKOI8(r rune) bool {
	return isUpperKOI8(r) || isLowerKOI8(r)
}
