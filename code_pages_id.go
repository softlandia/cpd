package cpd

const (
	// ASCII is the uint16 identifier with IANA name US-ASCII (MIME: US-ASCII).
	// ANSI X3.4-1986
	// Reference: RFC2046
	ASCII IDCodePage = 3

	// ISOLatinCyrillic is the MIB identifier with IANA name ISO_8859-5:1988 (MIME: ISO-8859-5).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatinCyrillic IDCodePage = 8

	// UTF8 is the uint16 identifier with IANA name UTF-8.
	//
	// rfc3629
	// Reference: RFC3629
	UTF8 IDCodePage = 106

	// Unicode is the uint16 identifier with IANA name ISO-10646-UCS-2.
	//
	// the 2-octet Basic Multilingual Plane, aka Unicode
	// this needs to specify network byte order: the standard
	// does not specify (it is a 16-bit integer space)
	Unicode IDCodePage = 1000

	// UnicodeASCII is the uint16 identifier with IANA name ISO-10646-UCS-Basic.
	//
	// ASCII subset of Unicode.  Basic Latin = collection 1
	// See ISO 10646, Appendix A
	UnicodeASCII IDCodePage = 1002

	// UTF7 is the uint16 identifier with IANA name UTF-7.
	//
	// rfc2152
	// Reference: RFC2152
	UTF7 IDCodePage = 1012

	// UTF16BE is the uint16 identifier with IANA name UTF-16BE.
	//
	// rfc2781
	// Reference: RFC2781
	UTF16BE IDCodePage = 1013

	// UTF16LE is the uint16 identifier with IANA name UTF-16LE.
	//
	// rfc2781
	// Reference: RFC2781
	UTF16LE IDCodePage = 1014

	// UTF32 is the uint16 identifier with IANA name UTF-32.
	//
	// https://www.unicode.org/unicode/reports/tr19/
	UTF32 IDCodePage = 1017

	// UTF32BE is the uint16 identifier with IANA name UTF-32BE.
	//
	// https://www.unicode.org/unicode/reports/tr19/
	UTF32BE IDCodePage = 1018

	// UTF32LE is the uint16 identifier with IANA name UTF-32LE.
	//
	// https://www.unicode.org/unicode/reports/tr19/
	UTF32LE IDCodePage = 1019

	// KOI8R is the uint16 identifier with IANA name KOI8-R (MIME: KOI8-R).
	//
	// rfc1489 , based on GOST-19768-74, ISO-6937/8,
	// INIS-Cyrillic, ISO-5427.
	// Reference: RFC1489
	KOI8R IDCodePage = 2084

	// CP866 is the uint16 identifier with IANA name IBM866.
	//
	// IBM NLDG Volume 2 (SE09-8002-03) August 1994
	CP866 IDCodePage = 2086

	// CP1251 is the uint16 identifier with IANA name windows-1251.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1251
	CP1251 IDCodePage = 2251

	// Windows1252 is the uint16 identifier with IANA name windows-1252.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1252
	Windows1252 IDCodePage = 2252
)
