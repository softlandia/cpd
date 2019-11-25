package cpd

//UTF-32LE

//первые 2 байта практически всегда больше вторых 2 байтов
//TODO лучше попробовать оценить количество 0x0 байтов по отношению к общему и если их много, то только тогла определять LE/BE
func runesMatchUTF32LE(d []byte, tbl *codePageTable) (counts int) {
	var (
		w1 int64
		w2 int64
	)
	for i := 0; i < len(d)-4; i += 4 {
		w1 = int64(d[i]) * int64(d[i+1])
		w2 = int64(d[i+2]) * int64(d[i+3])
		if w1 < w2 {
			counts = 0 //все первые должны быть больше, иначе это не UTF-32le
			break
		}
		counts++
	}
	return
}
