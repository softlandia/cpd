// file from "golang.org\x\text\encoding\internal\identifier" (c) golang autors
// contain identifier of code page
// IDCodePage implements interface String()

package cpd

import (
	"fmt"
	"strings"
)

//IDCodePage - index of code page
type IDCodePage uint16

func (i IDCodePage) String() string {
	return codePageName[i]
}

//itRuneMatch - return 1 if rune from this code page, 0 else
// function exist in every CodePage
//type itRuneMatch func(r rune, tbl *codePageTable) int

//runesMatch - return count of entry elements of data to code page
type runesMatch func(data []byte, tbl *codePageTable) MatchRes

type tableElement struct {
	code  rune //руна которая нас интересует, она присутствует в этой кодовой таблице как буква алфавита
	count int  //количество вхождений данной руны
}

// codePageTable - содержит основные (наиболее часто встречаемые) символы алфавита в данной кодировке
// первые 9 прописные, 2-я девятка заглавные
type codePageTable [19]tableElement

// MatchRes - итоговый критерий совпадения массива данных с кодовой страницей
// возможно в дальнейшем усложнится
type MatchRes struct {
	countMatch   int
	countCvPairs int
}

func (m MatchRes) String() string {
	return fmt.Sprintf("%d, %d", m.countMatch, m.countCvPairs)
}

// CodePage - содержит данные по конкретной кодовой странице
type CodePage struct {
	id       IDCodePage    //id of code page
	name     string        //name of code page
	MatchRes               //count of matching
	match    runesMatch    //method to calculate from input data count of entry to codepage
	table    codePageTable //table of main alphabet rune of this code page, contain [code, count]
}

func (o CodePage) String() string {
	return fmt.Sprintf("id: %s, countMatch: %d", o.id, o.countMatch)
}

// MatchingRunes - return string with rune/counts
func (o CodePage) MatchingRunes() string {
	var sb strings.Builder
	fmt.Fprint(&sb, "rune/counts: ")
	for i, e := range o.table {
		if i != 0 {
			fmt.Fprintf(&sb, "%x/%d, ", e.code, e.count)
		}
	}
	return sb.String()
}

//TCodepagesDic - type to store all supported code page
type TCodepagesDic map[IDCodePage]CodePage

//CodepageDic -
var CodepageDic = TCodepagesDic{
	ASCII: {ASCII, "ASCII", MatchRes{0, 0}, matchASCII,
		codePageTable{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}},

	CP866: {CP866, "CP866", MatchRes{0, 0}, match866,
		codePageTable{
			//first element serves as sign of absence
			{0, 0},
			//о          е		   а		  и			 н			т			с		  р			в
			{0xAE, 0}, {0xA5, 0}, {0xA0, 0}, {0xA8, 0}, {0xAD, 0}, {0xE2, 0}, {0xE1, 0}, {0xE0, 0}, {0xA2, 0},
			{0x8E, 0}, {0x85, 0}, {0x80, 0}, {0x88, 0}, {0x8D, 0}, {0x92, 0}, {0x91, 0}, {0x90, 0}, {0x82, 0}}},
	CP1251: {CP1251, "CP1251", MatchRes{0, 0}, match1251,
		codePageTable{
			{0, 0},
			//а		    и		   н		  с			 р			в		   л		  к			 я
			{0xE0, 0}, {0xE8, 0}, {0xED, 0}, {0xF1, 0}, {0xF0, 0}, {0xE2, 0}, {0xEB, 0}, {0xEA, 0}, {0xFF, 0},
			{0xC0, 0}, {0xC8, 0}, {0xCD, 0}, {0xD1, 0}, {0xD0, 0}, {0xC2, 0}, {0xCB, 0}, {0xCA, 0}, {0xDF, 0}}},
	KOI8R: {KOI8R, "KOI8-R", MatchRes{0, 0}, matchKOI8,
		codePageTable{
			//о		    а		   и		  т			 с			в		   л		  к			м
			{0, 0},
			{0xCF, 0}, {0xC1, 0}, {0xC9, 0}, {0xD4, 0}, {0xD3, 0}, {0xD7, 0}, {0xCC, 0}, {0xCB, 0}, {0xCD, 0},
			{0xEF, 0}, {0xE1, 0}, {0xE9, 0}, {0xF4, 0}, {0xF3, 0}, {0xF7, 0}, {0xEC, 0}, {0xEB, 0}, {0xED, 0}}},
	ISOLatinCyrillic: {ISOLatinCyrillic, "ISO-8859-5", MatchRes{0, 0}, matchISO88595,
		codePageTable{
			//о		    а		   и		  т			 с			в		   л		  к			е
			{0, 0},
			{0xDE, 0}, {0xD0, 0}, {0xD8, 0}, {0xE2, 0}, {0xE1, 0}, {0xD2, 0}, {0xDB, 0}, {0xDA, 0}, {0xD5, 0},
			{0xBF, 0}, {0xB0, 0}, {0xB8, 0}, {0xC2, 0}, {0xC1, 0}, {0xB2, 0}, {0xBB, 0}, {0xBA, 0}, {0xB5, 0}}},
	UTF8: {UTF8, "UTF-8", MatchRes{0, 0}, matchUTF8,
		codePageTable{
			{0, 0},
			//о           е				а		    и			 н			  т			   с			р			в
			{0xD0BE, 0}, {0xD0B5, 0}, {0xD0B0, 0}, {0xD0B8, 0}, {0xD0BD, 0}, {0xD182, 0}, {0xD181, 0}, {0xD180, 0}, {0xD0B2, 0},
			{0xD09E, 0}, {0xD095, 0}, {0xD090, 0}, {0xD098, 0}, {0xD0AD, 0}, {0xD0A2, 0}, {0xD0A1, 0}, {0xD0A0, 0}, {0xD092, 0}}},
	UTF16LE: {UTF16LE, "UTF-16LE", MatchRes{0, 0}, matchUTF16LE,
		codePageTable{
			{0, 0},
			//о           е				а		    и			 н			  т			   с			р			в
			{0x3E04, 0}, {0x3504, 0}, {0x1004, 0}, {0x3804, 0}, {0x3D04, 0}, {0x4204, 0}, {0x4104, 0}, {0x4004, 0}, {0x3204, 0},
			{0x1E04, 0}, {0x1504, 0}, {0x3004, 0}, {0x1804, 0}, {0x1D04, 0}, {0x2204, 0}, {0x2104, 0}, {0x2004, 0}, {0x1204, 0}}},
	UTF16BE: {UTF16BE, "UTF-16BE", MatchRes{0, 0}, matchUTF16BE,
		codePageTable{
			{0, 0},
			//о           е				а		    и			 н			  т			   с			р			в
			{0x043E, 0}, {0x0435, 0}, {0x0410, 0}, {0x0438, 0}, {0x043D, 0}, {0x0442, 0}, {0x0441, 0}, {0x0440, 0}, {0x0432, 0},
			{0x041E, 0}, {0x0415, 0}, {0x0430, 0}, {0x0418, 0}, {0x041D, 0}, {0x0422, 0}, {0x0421, 0}, {0x0420, 0}, {0x0412, 0}}},
	UTF32BE: {UTF32BE, "UTF-32BE", MatchRes{0, 0}, matchUTF32be,
		codePageTable{
			{0, 0},
			{0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0},
			{0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}}},
	UTF32LE: {UTF32LE, "UTF-32LE", MatchRes{0, 0}, matchUTF32le,
		codePageTable{
			{0, 0},
			{0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0},
			{0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}}},
}

//если кодировка файла определяется по BOM или по внутренней структуре,
//то счётчики совпадений будут заполнены не корректно
//обнуляем счётчики
func (o TCodepagesDic) clear() {
	for id, cp := range o {
		cp.countMatch = 0
		cp.table.clear()
		o[id] = cp
	}
}

//Match - return the id of code page to which the data best matches
func (o TCodepagesDic) Match(data []byte) (result IDCodePage) {
	result = ASCII
	maxCount := 0
	m := 0
	for id, cp := range o {
		cp.MatchRes = cp.match(data, &cp.table)
		o[id] = cp
		m = cp.countMatch + cp.countCvPairs
		if m > maxCount {
			maxCount = m
			result = id
		}
	}
	return result
}

//codePageName - string of code page name runesMatchUTF32LE
var codePageName = map[IDCodePage]string{
	ASCII:            "ASCII",
	ISOLatinCyrillic: "ISO-8859-5",
	CP866:            "CP866",
	CP1251:           "CP1251",
	UTF8:             "UTF-8",
	UTF16LE:          "UTF-16LE",
	UTF16BE:          "UTF-16BE",
	UTF32:            "UTF-32",
	KOI8R:            "KOI8-R",
	Unicode:          "Unicode",
	UTF7:             "UTF-7",
	UTF32LE:          "UTF-32LE",
	UTF32BE:          "UTF-32BE",
}

func matchASCII(b []byte, tbl *codePageTable) MatchRes {
	return MatchRes{0, 0}
}
