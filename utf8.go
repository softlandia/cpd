package cpd

import "encoding/binary"

//unit for UTF8

func runesMatchUTF8(data []byte, tbl *codePageTable) (counts int) {
	n := len(data)/2 - 1
	if n <= 0 {
		return
	}
	for i := 0; i < n; i += 2 {
		t := data[i : i+2]
		d := binary.BigEndian.Uint16(t)
		j := tbl.containsRune(rune(d))
		if j > 0 {
			(*tbl)[j].count++
			counts++
		}
	}
	return
}
