package cpd

//UTF-32LE

//вторые 2 байта практически всегда 0
//используемый признак не сработает если больше половины текста будет набрано символами с 4 значащими байтами, не представляю, что это за текст...
func runesMatchUTF32LE(d []byte, tbl *codePageTable) (zerroCounts int) {
	for i := 0; i < len(d)-4; i += 4 {
		if (int(d[i+2]) + int(d[i+3])) == 0 {
			zerroCounts++
		}
	}
	if zerroCounts*2 < len(d)/4 { //количество байтов в файле UTF-32 со значением 0 должно быть больше половины
		return 0
	}
	return
}
