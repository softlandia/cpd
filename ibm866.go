package cpd

//unit for ibm866

//tbl - массив структур (19 штук), 18 рабочих, нулевой элемент технический
//каждая структура состоит из искомого символа и счётчика
//для CP866 достаточно только вычисления счётчика по базовому алгоритму:
//подсчитываем случаи когда нужные нам буквы просто встречаются
func match866(data []byte, tbl *codePageTable) MatchRes {
	matches := 0
	for i := range data {
		j := tbl.index(rune(data[i])) //получили номер символа в таблице
		(*tbl)[j].count++
		if j > 0 {
			matches++
		}
	}
	return MatchRes{matches, 0}
}

const (
	cp866BeginUpperChar  = 0x80
	cp866StopUpperChar   = 0x9F
	cp866BeginLowerChar1 = 0xA0
	cp866StopLowerChar1  = 0xAF
	cp866BeginLowerChar2 = 0xE0
	cp866StopLowerChar2  = 0xEF
)

func isUpper866(r byte) bool {
	return (r >= cp866BeginUpperChar) && (r <= cp866StopUpperChar)
}

func isLower866(r byte) bool {
	return ((r >= cp866BeginLowerChar1) && (r <= cp866StopLowerChar1)) ||
		((r >= cp866BeginLowerChar2) && (r <= cp866StopLowerChar2))
}

func is866(r byte) bool {
	return isUpper866(r) || isLower866(r)
}
