package cpd

//unit for ibm866

//tbl - массив структур (19 штук), 18 рабочих, нулевой элемент технический
//каждая структура состоит из искомого символа и счётчика
//для CP866 достаточно только вычисления счётчика по базовому алгоритму:
//подсчитываем случаи когда нужные нам буквы просто встречаются
//TODO возможно стоит вынести (*tbl)[j].count++ из if поскольку даже если символ не найден, то увеличить счётчик с индексом 0 можно. зато читаемость улучшится
//тем более нулевой элемент таблицы мы ведь завели, собственно для этого
func runesMatch866(data []byte, tbl *codePageTable) (foundedCounts int) {
	for i := range data {
		j := tbl.containsRune(rune(data[i])) //получили номер символа в таблице
		(*tbl)[j].count++
		if j > 0 {
			foundedCounts++
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
