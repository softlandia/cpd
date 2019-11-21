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
			//counts++
		}
		if isUTF8(rune(d)) {
			counts++
		}
	}
	return
}

//тест на возможное начало байтов в UTF-8
//возвращает количество байтов из которых состоит код символа, возвращает служебное значение cp
//первый байт ВСЕГДА должен начинаться так:
//0xxx xxxx - здесь все символы ASCII 127
//110x xxxx - символ из 2-х байтов
//1110 xxxx - символ из 3-х байтов
//1111 0xxx - символ из 4-х байтов
func testUTF8bitPattern(b byte) (int, cp int32) {
	if (b & 0x80) == 0x00 { //b & 1000 0000 == 0x0
		cp = int32(b & 0x7F)
		return 1, cp
	}
	if (b & 0xE0) == 0xC0 { //b & 1110 0000 == 1100 0000
		cp = int32(b & 0x1F)
		return 2, cp
	}
	if (b & 0xF0) == 0xE0 { //b & 1111 0000 == 1110 0000
		cp = int32(b & 0x0F)
		return 3, cp
	}
	if (b & 0xF8) == 0xF0 { //b & 1111 1000 == 1111 0000
		cp = int32(b & 0x07)
		return 4, cp
	}
	return 0, 0
}

//ValidUTF8 - return true if imput slice contain true UTF-8
func ValidUTF8(data []byte) bool {
	m := len(data)
	if m <= 1 {
		return true
	}
	zerroByteCount := 0
	for i := 0; i < m-1; {
		if (data[i] == 0x0) && (data[i+1] != 0x0) {
			zerroByteCount++
		}
		n, cp := testUTF8bitPattern(data[i])
		if n == 0 {
			return false
		}
		i++
		var j int32 = 1
		for ; j < n; j++ {
			if (data[j] & 0xC0) != 0x80 {
				return false
			}
			cp = (cp << 6) | int32(data[j]&0x3F)
			i++
		}

		if (cp > 0x10FFFF) ||
			((cp >= 0xD800) && (cp <= 0xDFFF)) ||
			((cp <= 0x007F) && (n != 1)) ||
			((cp >= 0x0080) && (cp <= 0x07FF) && (n != 2)) ||
			((cp >= 0x0800) && (cp <= 0xFFFF) && (n != 3)) ||
			((cp >= 0x10000) && (cp <= 0x1FFFFF) && (n != 4)) {
			return false
		}
	}
	if float64(zerroByteCount)/float64(m) > 0.05 {
		return false
	}
	return true
}

const (
	cpUTF8StartUpperChar = 0xD090
	cpUTF8StopUpperChar  = 0xD0AF
	cpUTF8StartLowerChar = 0xD0B0
	cpUTF8StopLowerChar  = 0xD18F
)

func isUpperUTF8(r rune) bool {
	return (r >= cpUTF8StartUpperChar) && (r <= cpUTF8StopUpperChar)
}

func isLowerUTF8(r rune) bool {
	return (r >= cpUTF8StartLowerChar) && (r <= cpUTF8StopLowerChar)
}

func isUTF8(r rune) bool {
	return isUpperUTF8(r) || isLowerUTF8(r)
}
