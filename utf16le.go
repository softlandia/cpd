package cpd

import "encoding/binary"

//unit for UTF16LE

func runesMatchUTF16LE(data []byte, tbl *codePageTable) (counts int) {
	n := len(data)/2 - 1
	if n <= 0 {
		return
	}
	count04 := 0
	for i := 0; i < n; i += 2 {
		if data[i+1] == 0x04 {
			count04++
		}
		t := data[i : i+2]
		d := binary.BigEndian.Uint16(t)
		j := tbl.containsRune(rune(d))
		if j > 0 {
			(*tbl)[j].count++
		}
		if isUTF16LE(rune(d)) {
			counts++
		}
	}
	if count04 < counts {
		counts = count04
	}
	return
}

const (
	cpUTF16LEStartUpperChar = 0x1004
	cpUTF16LEStopUpperChar  = 0x2F04
	cpUTF16LEStartLowerChar = 0x3004
	cpUTF16LEStopLowerChar  = 0x4F04
)

func isUpperUTF16LE(r rune) bool {
	return (r >= cpUTF16LEStartUpperChar) && (r <= cpUTF16LEStopUpperChar)
}

func isLowerUTF16LE(r rune) bool {
	return (r >= cpUTF16LEStartLowerChar) && (r <= cpUTF16LEStopLowerChar)
}

func isUTF16LE(r rune) bool {
	return isUpperUTF16LE(r) || isLowerUTF16LE(r)
}
