package cpd

import "unicode"

//unit for ISO-8859-5

func matchISO88595(d []byte, tbl *codePageTable) MatchRes {
	for i := 0; i < len(d); i++ {
		if isISO88595(rune(d[i])) {
			upper := lu88595(d[i])
			j := tbl.index(rune(d[i]))
			(*tbl)[j].count++
			for i++; (i < len(d)) && isISO88595(rune(d[i])); i++ {
				if upper >= lu88595(d[i]) {
					j = tbl.index(rune(d[i]))
					(*tbl)[j].count++
				}
			}
		}
	}
	return MatchRes{tbl.founded(), 0}
}

func runesMatchISO88595_2(data []byte, tbl *codePageTable) (counts int) {
	for i := range data {
		if i < 2 {
			continue
		}
		//case " Us" - separator_UPPER_symbol
		if unicode.IsPunct(rune(data[i-2])) && isUpperISO88595(rune(data[i-1])) {
			j := tbl.index(rune(data[i]))
			if j > 0 {
				(*tbl)[j].count++
				counts++
				continue
			}
		}
		if isISO88595(rune(data[i-1])) {
			j := tbl.index(rune(data[i]))
			if j > 0 {
				(*tbl)[j].count++
				counts++
			}
		}
	}
	return
}

const (
	cpISO88595BeginUpperChar = 0xB0
	cpISO88595StopUpperChar  = 0xCF
	cpISO88595BeginLowerChar = 0xD0
	cpISO88595StopLowerChar  = 0xEF
)

func lu88595(r byte) (res int) {
	if isUpperISO88595(rune(r)) {
		res = 1
	}
	return
}

func isUpperISO88595(r rune) bool {
	return (r >= cpKOI8BeginUpperChar) && (r <= cpKOI8StopUpperChar)
}

func isLowerISO88595(r rune) bool {
	return (r >= cpKOI8BeginLowerChar) && (r <= cpKOI8StopLowerChar)
}

func isISO88595(r rune) bool {
	return isUpperISO88595(r) || isLowerISO88595(r)
}
