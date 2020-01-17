package cpd

import (
	"bytes"
	"fmt"
	"strings"
)

// IDCodePage - index of code page
// implements interface String()
type IDCodePage uint16

func (i IDCodePage) String() string {
	//return codePageName[i]
	return CodepageDic[i].name
}

//StringHasBom - return true if input string has BOM prefix
func (i IDCodePage) StringHasBom(s string) bool {
	if len(CodepageDic[i].Boms) == 0 {
		return false
	}
	return bytes.HasPrefix([]byte(s), CodepageDic[i].Boms)
}

//DeleteBom - return string without prefix bom bytes
func (i IDCodePage) DeleteBom(s string) (res string) {
	res = s
	if i.StringHasBom(s) {
		res = res[len(CodepageDic[i].Boms):]
	}
	return res
}

// matcher - return struct MatchRes - two criterion
// this function must be realised in each code page
type matcher func(data []byte, tbl *codePageTable) MatchRes

type tableElement struct {
	code  rune //rune (letter) of the alphabet that interests us
	count int  //the number of these runes found in the text
}

// codePageTable - stores 9 letters, we will look for them in the text
// element with index 0 for the case of non-location
// first 9 elements lowercase, second 9 elements uppercase
type codePageTable [19]tableElement

// MatchRes - result criteria
// countMatch - the number of letters founded in text
// countCvPairs - then number of pairs consonans+vowels
type MatchRes struct {
	countMatch   int
	countCvPairs int
}

func (m MatchRes) String() string {
	return fmt.Sprintf("%d, %d", m.countMatch, m.countCvPairs)
}

// CodePage - realize code page
type CodePage struct {
	id       IDCodePage    //id of code page
	name     string        //name of code page
	MatchRes               //count of matching
	match    matcher       //method for calculating the criteria for the proximity of input data to this code page
	Boms     []byte        //default BOM for this codepage
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

// TCodepagesDic - type to store all supported code page
type TCodepagesDic map[IDCodePage]CodePage

//CodepageDic - map of all codepage
var CodepageDic = TCodepagesDic{
	ASCII: {ASCII, "ASCII", MatchRes{0, 0}, matchASCII, []byte{},
		codePageTable{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}},

	CP866: {CP866, "CP866", MatchRes{0, 0}, match866, []byte{},
		codePageTable{
			//first element serves as sign of absence
			{0, 0},
			//о          е		   а		  и			 н			т			с		  р			в
			{0xAE, 0}, {0xA5, 0}, {0xA0, 0}, {0xA8, 0}, {0xAD, 0}, {0xE2, 0}, {0xE1, 0}, {0xE0, 0}, {0xA2, 0},
			{0x8E, 0}, {0x85, 0}, {0x80, 0}, {0x88, 0}, {0x8D, 0}, {0x92, 0}, {0x91, 0}, {0x90, 0}, {0x82, 0}}},
	CP1251: {CP1251, "CP1251", MatchRes{0, 0}, match1251, []byte{},
		codePageTable{
			{0, 0},
			//а		    и		   н		  с			 р			в		   л		  к			 я
			{0xE0, 0}, {0xE8, 0}, {0xED, 0}, {0xF1, 0}, {0xF0, 0}, {0xE2, 0}, {0xEB, 0}, {0xEA, 0}, {0xFF, 0},
			{0xC0, 0}, {0xC8, 0}, {0xCD, 0}, {0xD1, 0}, {0xD0, 0}, {0xC2, 0}, {0xCB, 0}, {0xCA, 0}, {0xDF, 0}}},
	KOI8R: {KOI8R, "KOI8-R", MatchRes{0, 0}, matchKOI8, []byte{},
		codePageTable{
			//о		    а		   и		  т			 с			в		   л		  к			м
			{0, 0},
			{0xCF, 0}, {0xC1, 0}, {0xC9, 0}, {0xD4, 0}, {0xD3, 0}, {0xD7, 0}, {0xCC, 0}, {0xCB, 0}, {0xCD, 0},
			{0xEF, 0}, {0xE1, 0}, {0xE9, 0}, {0xF4, 0}, {0xF3, 0}, {0xF7, 0}, {0xEC, 0}, {0xEB, 0}, {0xED, 0}}},
	ISOLatinCyrillic: {ISOLatinCyrillic, "ISO-8859-5", MatchRes{0, 0}, matchISO88595, []byte{},
		codePageTable{
			//о		    а		   и		  т			 с			в		   л		  к			е
			{0, 0},
			{0xDE, 0}, {0xD0, 0}, {0xD8, 0}, {0xE2, 0}, {0xE1, 0}, {0xD2, 0}, {0xDB, 0}, {0xDA, 0}, {0xD5, 0},
			{0xBF, 0}, {0xB0, 0}, {0xB8, 0}, {0xC2, 0}, {0xC1, 0}, {0xB2, 0}, {0xBB, 0}, {0xBA, 0}, {0xB5, 0}}},
	UTF8: {UTF8, "UTF-8", MatchRes{0, 0}, matchUTF8, []byte{0xef, 0xbb, 0xbf},
		codePageTable{
			{0, 0},
			//о           е				а		    и			 н			  т			   с			р			в
			{0xD0BE, 0}, {0xD0B5, 0}, {0xD0B0, 0}, {0xD0B8, 0}, {0xD0BD, 0}, {0xD182, 0}, {0xD181, 0}, {0xD180, 0}, {0xD0B2, 0},
			{0xD09E, 0}, {0xD095, 0}, {0xD090, 0}, {0xD098, 0}, {0xD0AD, 0}, {0xD0A2, 0}, {0xD0A1, 0}, {0xD0A0, 0}, {0xD092, 0}}},
	UTF16LE: {UTF16LE, "UTF-16LE", MatchRes{0, 0}, matchUTF16LE, []byte{0xff, 0xfe},
		codePageTable{
			{0, 0},
			//о           е				а		    и			 н			  т			   с			р			в
			{0x3E04, 0}, {0x3504, 0}, {0x1004, 0}, {0x3804, 0}, {0x3D04, 0}, {0x4204, 0}, {0x4104, 0}, {0x4004, 0}, {0x3204, 0},
			{0x1E04, 0}, {0x1504, 0}, {0x3004, 0}, {0x1804, 0}, {0x1D04, 0}, {0x2204, 0}, {0x2104, 0}, {0x2004, 0}, {0x1204, 0}}},
	UTF16BE: {UTF16BE, "UTF-16BE", MatchRes{0, 0}, matchUTF16BE, []byte{0xfe, 0xff},
		codePageTable{
			{0, 0},
			//о           е				а		    и			 н			  т			   с			р			в
			{0x043E, 0}, {0x0435, 0}, {0x0410, 0}, {0x0438, 0}, {0x043D, 0}, {0x0442, 0}, {0x0441, 0}, {0x0440, 0}, {0x0432, 0},
			{0x041E, 0}, {0x0415, 0}, {0x0430, 0}, {0x0418, 0}, {0x041D, 0}, {0x0422, 0}, {0x0421, 0}, {0x0420, 0}, {0x0412, 0}}},
	UTF32BE: {UTF32BE, "UTF-32BE", MatchRes{0, 0}, matchUTF32be, []byte{0x00, 0x00, 0xfe, 0xff},
		codePageTable{
			{0, 0},
			{0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0},
			{0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}}},
	UTF32LE: {UTF32LE, "UTF-32LE", MatchRes{0, 0}, matchUTF32le, []byte{0xff, 0xfe, 0x00, 0x00},
		codePageTable{
			{0, 0},
			{0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0},
			{0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}, {0x0, 0}}},
}

//befor detecting of code page need clear all counts
//this not for correct run, this need only if we want get correct statistic
func (o TCodepagesDic) clear() {
	for id, cp := range o {
		cp.MatchRes = MatchRes{0, 0}
		cp.table.clear()
		o[id] = cp
	}
}

//Match - return the id of code page to which the data best matches
//TODO большинству матчеров требуется более 2х символов, надо проверить на минимальную длину
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

//foo function,
func matchASCII(b []byte, tbl *codePageTable) MatchRes {
	return MatchRes{0, 0}
}

/*
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
*/
