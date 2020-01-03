package cpd

import "encoding/binary"

//unit for UTF16LE

func matchUTF16LE(data []byte, tbl *codePageTable) MatchRes {
	matches := 0
	n := len(data)/2 - 1
	if n <= 0 {
		return MatchRes{0, 0}
	}
	count04 := 0
	for i := 0; i < n; i += 2 {
		if data[i+1] == 0x04 {
			count04++
		}
		t := data[i : i+2]
		d := binary.BigEndian.Uint16(t)
		j := tbl.index(rune(d))
		if j > 0 {
			(*tbl)[j].count++
		}
		if isUTF16LE(rune(d)) {
			matches++
		}
	}
	if count04 < matches {
		matches = count04
	}
	return MatchRes{matches, 0}
}

const (
	cpUTF16leBeginUpperChar = 0x1004
	cpUTF16leStopUpperChar  = 0x2F04
	cpUTF16leBeginLowerChar = 0x3004
	cpUTF16leStopLowerChar  = 0x4F04
)

func isUpperUTF16LE(r rune) bool {
	return (r >= cpUTF16leBeginUpperChar) && (r <= cpUTF16leStopUpperChar)
}

func isLowerUTF16LE(r rune) bool {
	return (r >= cpUTF16leBeginLowerChar) && (r <= cpUTF16leStopLowerChar)
}

func isUTF16LE(r rune) bool {
	return isUpperUTF16LE(r) || isLowerUTF16LE(r)
}
