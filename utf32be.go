package cpd

//UTF-32BE

//первые 2 байта практически всегда меньше вторых 2 байтов
func runesMatchUTF32BE(d []byte, tbl *codePageTable) (zerroCounts int) {
	for i := 0; i < len(d)-4; i += 4 {
		if (int(d[i]) + int(d[i+1])) == 0 {
			zerroCounts++
		}
	}
	if zerroCounts*2 < len(d)/4 { //количество байтов в файле UTF-32 со значением 0 должно быть больше половины
		return 0
	}
	return

}
