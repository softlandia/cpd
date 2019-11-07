package cpd

//unit for ibm866

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
