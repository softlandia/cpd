package cpd

//UTF-32BE

//первые 2 байта практически всегда меньше вторых 2 байтов
func runesMatchUTF32BE(d []byte, tbl *codePageTable) (counts int) {
	var (
		w1 int64
		w2 int64
	)
	for i := 0; i < len(d)-4; i += 4 {
		w1 = int64(d[i]) * int64(d[i+1])
		w2 = int64(d[i+2]) * int64(d[i+3])
		if w1 > w2 {
			counts = 0
			break
		}
		counts++
	}
	return
}
