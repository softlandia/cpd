package cpd

//unit for ibm866

//tbl - массив структур (19 штук), 18 рабочих, нулевой элемент технический
//каждая структура состоит из искомого символа и счётчика
//для CP866 достаточно только вычисления счётчика по базовому алгоритму:
//подсчитываем случи когда нужные нам буквы встречаются подряд
func runesMatch866(data []byte, tbl *codePageTable) (counts int) {
	for i := range data {
		if i == 0 {
			continue
		}
		if tbl.containsRune(rune(data[i-1])) > 0 {
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
	cp866StartUpperChar  = 0x80
	cp866StopUpperChar   = 0x9F
	cp866StartLowerChar1 = 0xA0
	cp866StopLowerChar1  = 0xAF
	cp866StartLowerChar2 = 0xE0
	cp866StopLowerChar2  = 0xEF
)

func isUpper866(r rune) bool {
	return (r >= cp866StartUpperChar) && (r <= cp866StopUpperChar)
}

func isLower866(r rune) bool {
	return ((r >= cp866StartLowerChar1) && (r <= cp866StopLowerChar1)) ||
		((r >= cp866StartLowerChar2) && (r <= cp866StopLowerChar2))
}

func is866(r rune) bool {
	return isUpper866(r) || isLower866(r)
}
