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
type itRuneMatch func(r rune, tbl *codePageTable) int

//runesMatch - return count of entry elements of data to code page
type runesMatch func(data []byte, tbl *codePageTable) int

type tableElement struct {
	code  rune //руна которая нас интересует, она присутствует в этой кодовой таблице как буква алфавита
	count int  //количество вхождений данной руны
}

//codePageTable - содержит основные (наиболее часто встречаемые) символы алфавита в данной кодировке
//первые 9 прописные, 2-я девятка заглавные
type codePageTable [19]tableElement

//MatchRes - итоговый критерий совпадения массива данных с кодовой страницей
type MatchRes struct {
	countMatch int
}

//CodePage - содержит данные по конкретной кодовой странице
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

//MatchingRunes - return string with rune/counts
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

//TCodePages - type for store all code page
type TCodePages []CodePage

//DeepMach -
func (o *TCodePages) DeepMach(data []byte) IDCodePage {
	return ASCII
}

//Match - return IDCodePage
//simple calculate count entry data runes in standart code page table
func (o TCodePages) Match(data []byte) (result IDCodePage) {
	result = ASCII
	maxCount := 0
	for i, cp := range o {
		o[i].countMatch = cp.match(data, &o[i].table)
		if o[i].countMatch > maxCount {
			maxCount = o[i].countMatch
			result = cp.id
		}
	}
	return result
}

//CodePages - slice of code pages
var CodePages = TCodePages{
	{ASCII, "ASCII", MatchRes{0}, runesMatchASCII,
		codePageTable{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}},
	{IBM866, "IBM866", MatchRes{0}, runesMatch866,
		codePageTable{
			//first element serves as sign of absence
			{0, 0},
			//о          е		   а		  и			 н			т			с		  р			в
			{0xAE, 0}, {0xA5, 0}, {0xA0, 0}, {0xA8, 0}, {0xAD, 0}, {0xE2, 0}, {0xE1, 0}, {0xE0, 0}, {0xA2, 0},
			{0x8E, 0}, {0x85, 0}, {0x80, 0}, {0x88, 0}, {0x8D, 0}, {0x92, 0}, {0x91, 0}, {0x90, 0}, {0x82, 0}}},
	{UTF8, "UTF8", MatchRes{0}, runesMatchUTF8,
		codePageTable{
			{0, 0},
			//о           е				а		    и			 н			  т			   с			р			в
			{0xD0BE, 0}, {0xD0B5, 0}, {0xD0B0, 0}, {0xD0B8, 0}, {0xD0BD, 0}, {0xD182, 0}, {0xD181, 0}, {0xD180, 0}, {0xD0B2, 0},
			{0xD09E, 0}, {0xD095, 0}, {0xD090, 0}, {0xD098, 0}, {0xD0AD, 0}, {0xD0A2, 0}, {0xD0A1, 0}, {0xD0A0, 0}, {0xD092, 0}}},
	{Windows1251, "Windows1251", MatchRes{0}, runesMatch1251,
		codePageTable{
			{0, 0},
			//а		    и		   н		  с			 р			в		   л		  к			в
			{0xE0, 0}, {0xE8, 0}, {0xED, 0}, {0xF1, 0}, {0xF0, 0}, {0xE2, 0}, {0xEB, 0}, {0xEA, 0}, {0xE2, 0},
			{0xC0, 0}, {0xC8, 0}, {0xCD, 0}, {0xD1, 0}, {0xD0, 0}, {0xC2, 0}, {0xCB, 0}, {0xCA, 0}, {0xC2, 0}}},
	{KOI8R, "KOI8R", MatchRes{0}, runesMatchKOI8,
		codePageTable{
			//о		    а		   и		  т			 с			в		   л		  к			в
			{0, 0},
			{0xCF, 0}, {0xC1, 0}, {0xC9, 0}, {0xD4, 0}, {0xD3, 0}, {0xD7, 0}, {0xCC, 0}, {0xCB, 0}, {0xD7, 0},
			{0xEF, 0}, {0xE1, 0}, {0xE9, 0}, {0xF4, 0}, {0xF3, 0}, {0xF7, 0}, {0xEC, 0}, {0xEB, 0}, {0xF7, 0}}},
}

//codePageName - string of code page name
var codePageName = map[IDCodePage]string{
	ASCII:              "ASCII",
	IBM866:             "IBM866",
	Windows1251:        "Windows1251",
	UTF8:               "UTF8",
	UTF16:              "UTF16",
	UTF16LE:            "UTF16LE",
	UTF16BE:            "UTF16BE",
	UTF32:              "UTF32",
	KOI8R:              "KOI8R",
	ISO5427Cyrillic:    "ISO5427Cyrillic",
	ISO51INISCyrillic:  "ISO51INISCyrillic",
	ISO111ECMACyrillic: "ISO111ECMACyrillic",
	ISO153GOST1976874:  "ISO153GOST1976874",
	Unicode:            "Unicode",
}
