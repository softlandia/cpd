package cpd

import "encoding/binary"

//unit for UTF16BE

func matchUTF16BE(data []byte, tbl *codePageTable) MatchRes {
	matches := 0
	n := len(data)/2 - 1
	if n <= 0 {
		return MatchRes{0, 0}
	}
	count04 := 0
	for i := 0; i < n; i += 2 {
		if data[i] == 0x04 {
			count04++
		}
		t := data[i : i+2]
		d := binary.BigEndian.Uint16(t)
		j := tbl.index(rune(d))
		if j > 0 {
			(*tbl)[j].count++
		}
		if isUTF16BE(rune(d)) {
			matches++
		}
	}
	if count04 < matches {
		matches = count04
	}
	return MatchRes{matches, 0}
}

const (
	cpUTF16beBeginUpperChar = 0x0410
	cpUTF16BEStopUpperChar  = 0x042F
	cpUTF16beBeginLowerChar = 0x0430
	cpUTF16BEStopLowerChar  = 0x044F
)

func isUpperUTF16BE(r rune) bool {
	return (r >= cpUTF16beBeginUpperChar) && (r <= cpUTF16BEStopUpperChar)
}

func isLowerUTF16BE(r rune) bool {
	return (r >= cpUTF16beBeginLowerChar) && (r <= cpUTF16BEStopLowerChar)
}

func isUTF16BE(r rune) bool {
	return isUpperUTF16BE(r) || isLowerUTF16BE(r)
}
