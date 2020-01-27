package cpd

import (
	"encoding/binary"

	"github.com/softlandia/xlib"
)

//unit for UTF16LE

//русские буквы в UTF16 имеют уникальные номера
//определять кодировку UTF16 (как LE так и BE) нужно по внутреннему устройству, не по кодам русских букв
//проверка на BOM уже выполнена, в принимаемом массиве не BOM символов

// matchUTF16le - функция вычисляет общий критерий для кодировки UTF16LE
func matchUTF16le(b []byte, tbl *cpTable) MatchRes {
	n := len(b)/2 - 1
	if n <= 0 {
		return MatchRes{0, 0}
	}
	//два критерия используется
	//первый количество найденных русских букв
	//второй количество найденных 0x00
	//решающим является максимальный
	return MatchRes{xlib.Max(matchUTF16leRu(b, tbl), matchUTF16leZerro(b)), 0}
}

// matchUTF16leZerro - вычисляет критерий по количеству нулевых байтов, текст набранный латинскими символами в колировке UTF16le будет вторым символом иметь 0x00
func matchUTF16leZerro(b []byte) int {
	zerroCount := 0
	n := len(b)/2 - 1
	for i := 0; i < n; i += 2 {
		if b[i+1] == 0x00 {
			zerroCount++
		}
	}
	return zerroCount
}

// matchUTF16leRu - вычисляет критерий по количеству русских букв
// tbl *codePageTable - передаётся не для нахождения кодировки, а для заполнения встречаемости популярных русских букв
func matchUTF16leRu(b []byte, tbl *cpTable) int {
	matches := 0
	count04 := 0
	n := len(b)/2 - 1
	for i := 0; i < n; i += 2 {
		//second byte of russian char is 0x04
		if b[i+1] == 0x04 {
			count04++
		}
		t := b[i : i+2]
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
	return matches
}

/*func matchUTF16leFirstGreateSecond(b []byte) int {
	count := 0
	n := len(b)/2 - 1
	for i := 0; i < n; i += 2 {
		//second byte of russian char is 0x04
		if b[i] > b[i+1] {
			count++
		}
	}
	return count
}*/

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
