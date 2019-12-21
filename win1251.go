package cpd

import "unicode"

//unit for windows1251

func runesMatch1251(data []byte, tbl *codePageTable) (counts int) {
	for i := range data {
		if i < 1 {
			continue
		}
		//case "ab" - counts only if symbols are arranged in pairs
		if is1251(rune(data[i-1])) {
			j := tbl.containsRune(rune(data[i]))
			if j > 0 {
				(*tbl)[j].count++
				counts++
			}
			continue
		}
		if i < 2 {
			continue
		}
		//case " Us" separator_UPPER_lower
		//IsPunct -
		if (unicode.IsPunct(rune(data[i-2])) || unicode.IsSpace(rune(data[i-2]))) && isUpper1251(rune(data[i-1])) {
			j := tbl.containsRune(rune(data[i]))
			if (j > 0) && (isLower1251(rune(data[i]))) {
				(*tbl)[j].count++
				counts++
				continue
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
