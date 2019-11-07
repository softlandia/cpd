package cpd

import "unicode"

//unit for windows1251

//TODO: нужно отличить от KOI-8r
func runesMatch1251(data []byte, tbl *codePageTable) (counts int) {
	for i := range data {
		if i < 2 {
			continue
		}
		//case " Us" - separator_UPPER_symbol
		if unicode.IsPunct(rune(data[i-2])) && isUpper1251(rune(data[i-1])) {
			j := tbl.containsRune(rune(data[i]))
			if j > 0 {
				(*tbl)[j].count++
				counts++
				continue
			}
		}
		//case "ab" - counts only if symbols are arranged in pairs
		if is1251(rune(data[i-1])) {
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
	cp1251StartUpperChar = 0xC0
	cp1251StopUpperChar  = 0xDF
	cp1251StartLowerChar = 0xE0
	cp1251StopLowerChar  = 0xFF
)

func isUpper1251(r rune) bool {
	return (r >= cp1251StartUpperChar) && (r <= cp1251StopUpperChar)
}

func isLower1251(r rune) bool {
	return (r >= cp1251StartLowerChar) && (r <= cp1251StopLowerChar)
}

func is1251(r rune) bool {
	return isUpper1251(r) || isLower1251(r)
}
