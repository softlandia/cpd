package cpd

import "unicode"

//unit for ISO-8859-5

func runesMatchISO88595(data []byte, tbl *codePageTable) (counts int) {
	for i := range data {
		if i < 2 {
			continue
		}
		//case " Us" - separator_UPPER_symbol
		if unicode.IsPunct(rune(data[i-2])) && isUpperISO88595(rune(data[i-1])) {
			j := tbl.containsRune(rune(data[i]))
			if j > 0 {
				(*tbl)[j].count++
				counts++
				continue
			}
		}
		if isISO88595(rune(data[i-1])) {
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
	cpISO88595StartUpperChar = 0xB0
	cpISO88595StopUpperChar  = 0xCF
	cpISO88595StartLowerChar = 0xD0
	cpISO88595StopLowerChar  = 0xEF
)

func isUpperISO88595(r rune) bool {
	return (r >= cpKOI8StartUpperChar) && (r <= cpKOI8StopUpperChar)
}

func isLowerISO88595(r rune) bool {
	return (r >= cpKOI8StartLowerChar) && (r <= cpKOI8StopLowerChar)
}

func isISO88595(r rune) bool {
	return isUpperKOI8(r) || isLowerKOI8(r)
}
