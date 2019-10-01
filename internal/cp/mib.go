// file from "golang.org\x\text\encoding\internal\identifier" (c) golang autors
// contain identifier of code page

package cp

const (
	// ASCII is the uint16 identifier with IANA name US-ASCII (MIME: US-ASCII).
	//
	// ANSI X3.4-1986
	// Reference: RFC2046
	ASCII uint16 = 3

	// ISOLatin1 is the uint16 identifier with IANA name ISO_8859-1:1987 (MIME: ISO-8859-1).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatin1 uint16 = 4

	// ISOLatin2 is the uint16 identifier with IANA name ISO_8859-2:1987 (MIME: ISO-8859-2).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatin2 uint16 = 5

	// ISOLatin3 is the uint16 identifier with IANA name ISO_8859-3:1988 (MIME: ISO-8859-3).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatin3 uint16 = 6

	// ISOLatin4 is the uint16 identifier with IANA name ISO_8859-4:1988 (MIME: ISO-8859-4).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatin4 uint16 = 7

	// ISOLatinCyrillic is the uint16 identifier with IANA name ISO_8859-5:1988 (MIME: ISO-8859-5).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatinCyrillic uint16 = 8

	// ISOLatinArabic is the uint16 identifier with IANA name ISO_8859-6:1987 (MIME: ISO-8859-6).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatinArabic uint16 = 9

	// ISOLatinGreek is the uint16 identifier with IANA name ISO_8859-7:1987 (MIME: ISO-8859-7).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1947
	// Reference: RFC1345
	ISOLatinGreek uint16 = 10

	// ISOLatinHebrew is the uint16 identifier with IANA name ISO_8859-8:1988 (MIME: ISO-8859-8).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatinHebrew uint16 = 11

	// ISOLatin5 is the uint16 identifier with IANA name ISO_8859-9:1989 (MIME: ISO-8859-9).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatin5 uint16 = 12

	// ISOLatin6 is the uint16 identifier with IANA name ISO-8859-10 (MIME: ISO-8859-10).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOLatin6 uint16 = 13

	// ISOTextComm is the uint16 identifier with IANA name ISO_6937-2-add.
	//
	// ISO-IR: International Register of Escape Sequences and ISO 6937-2:1983
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISOTextComm uint16 = 14

	// HalfWidthKatakana is the uint16 identifier with IANA name JIS_X0201.
	//
	// JIS X 0201-1976.   One byte only, this is equivalent to
	// JIS/Roman (similar to ASCII) plus eight-bit half-width
	// Katakana
	// Reference: RFC1345
	HalfWidthKatakana uint16 = 15

	// JISEncoding is the uint16 identifier with IANA name JIS_Encoding.
	//
	// JIS X 0202-1991.  Uses ISO 2022 escape sequences to
	// shift code sets as documented in JIS X 0202-1991.
	JISEncoding uint16 = 16

	// ShiftJIS is the uint16 identifier with IANA name Shift_JIS (MIME: Shift_JIS).
	//
	// This charset is an extension of csHalfWidthKatakana by
	// adding graphic characters in JIS X 0208.  The CCS's are
	// JIS X0201:1997 and JIS X0208:1997.  The
	// complete definition is shown in Appendix 1 of JIS
	// X0208:1997.
	// This charset can be used for the top-level media type "text".
	ShiftJIS uint16 = 17

	// EUCPkdFmtJapanese is the uint16 identifier with IANA name Extended_UNIX_Code_Packed_Format_for_Japanese (MIME: EUC-JP).
	//
	// Standardized by OSF, UNIX International, and UNIX Systems
	// Laboratories Pacific.  Uses ISO 2022 rules to select
	// code set 0: US-ASCII (a single 7-bit byte set)
	// code set 1: JIS X0208-1990 (a double 8-bit byte set)
	// restricted to A0-FF in both bytes
	// code set 2: Half Width Katakana (a single 7-bit byte set)
	// requiring SS2 as the character prefix
	// code set 3: JIS X0212-1990 (a double 7-bit byte set)
	// restricted to A0-FF in both bytes
	// requiring SS3 as the character prefix
	EUCPkdFmtJapanese uint16 = 18

	// EUCFixWidJapanese is the uint16 identifier with IANA name Extended_UNIX_Code_Fixed_Width_for_Japanese.
	//
	// Used in Japan.  Each character is 2 octets.
	// code set 0: US-ASCII (a single 7-bit byte set)
	// 1st byte = 00
	// 2nd byte = 20-7E
	// code set 1: JIS X0208-1990 (a double 7-bit byte set)
	// restricted  to A0-FF in both bytes
	// code set 2: Half Width Katakana (a single 7-bit byte set)
	// 1st byte = 00
	// 2nd byte = A0-FF
	// code set 3: JIS X0212-1990 (a double 7-bit byte set)
	// restricted to A0-FF in
	// the first byte
	// and 21-7E in the second byte
	EUCFixWidJapanese uint16 = 19

	// ISO4UnitedKingdom is the uint16 identifier with IANA name BS_4730.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO4UnitedKingdom uint16 = 20

	// ISO11SwedishForNames is the uint16 identifier with IANA name SEN_850200_C.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO11SwedishForNames uint16 = 21

	// ISO15Italian is the uint16 identifier with IANA name IT.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO15Italian uint16 = 22

	// ISO17Spanish is the uint16 identifier with IANA name ES.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO17Spanish uint16 = 23

	// ISO21German is the uint16 identifier with IANA name DIN_66003.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO21German uint16 = 24

	// ISO60Norwegian1 is the uint16 identifier with IANA name NS_4551-1.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO60Norwegian1 uint16 = 25

	// ISO69French is the uint16 identifier with IANA name NF_Z_62-010.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO69French uint16 = 26

	// ISO10646UTF1 is the uint16 identifier with IANA name ISO-10646-UTF-1.
	//
	// Universal Transfer Format (1), this is the multibyte
	// encoding, that subsets ASCII-7. It does not have byte
	// ordering issues.
	ISO10646UTF1 uint16 = 27

	// ISO646basic1983 is the uint16 identifier with IANA name ISO_646.basic:1983.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO646basic1983 uint16 = 28

	// INVARIANT is the uint16 identifier with IANA name INVARIANT.
	//
	// Reference: RFC1345
	INVARIANT uint16 = 29

	// ISO2IntlRefVersion is the uint16 identifier with IANA name ISO_646.irv:1983.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO2IntlRefVersion uint16 = 30

	// NATSSEFI is the uint16 identifier with IANA name NATS-SEFI.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	NATSSEFI uint16 = 31

	// NATSSEFIADD is the uint16 identifier with IANA name NATS-SEFI-ADD.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	NATSSEFIADD uint16 = 32

	// NATSDANO is the uint16 identifier with IANA name NATS-DANO.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	NATSDANO uint16 = 33

	// NATSDANOADD is the uint16 identifier with IANA name NATS-DANO-ADD.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	NATSDANOADD uint16 = 34

	// ISO10Swedish is the uint16 identifier with IANA name SEN_850200_B.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO10Swedish uint16 = 35

	// KSC56011987 is the uint16 identifier with IANA name KS_C_5601-1987.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	KSC56011987 uint16 = 36

	// ISO2022KR is the uint16 identifier with IANA name ISO-2022-KR (MIME: ISO-2022-KR).
	//
	// rfc1557 (see also KS_C_5601-1987)
	// Reference: RFC1557
	ISO2022KR uint16 = 37

	// EUCKR is the uint16 identifier with IANA name EUC-KR (MIME: EUC-KR).
	//
	// rfc1557 (see also KS_C_5861-1992)
	// Reference: RFC1557
	EUCKR uint16 = 38

	// ISO2022JP is the uint16 identifier with IANA name ISO-2022-JP (MIME: ISO-2022-JP).
	//
	// rfc1468 (see also rfc2237 )
	// Reference: RFC1468
	ISO2022JP uint16 = 39

	// ISO2022JP2 is the uint16 identifier with IANA name ISO-2022-JP-2 (MIME: ISO-2022-JP-2).
	//
	// rfc1554
	// Reference: RFC1554
	ISO2022JP2 uint16 = 40

	// ISO13JISC6220jp is the uint16 identifier with IANA name JIS_C6220-1969-jp.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO13JISC6220jp uint16 = 41

	// ISO14JISC6220ro is the uint16 identifier with IANA name JIS_C6220-1969-ro.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO14JISC6220ro uint16 = 42

	// ISO16Portuguese is the uint16 identifier with IANA name PT.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO16Portuguese uint16 = 43

	// ISO18Greek7Old is the uint16 identifier with IANA name greek7-old.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO18Greek7Old uint16 = 44

	// ISO19LatinGreek is the uint16 identifier with IANA name latin-greek.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO19LatinGreek uint16 = 45

	// ISO25French is the uint16 identifier with IANA name NF_Z_62-010_(1973).
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO25French uint16 = 46

	// ISO27LatinGreek1 is the uint16 identifier with IANA name Latin-greek-1.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO27LatinGreek1 uint16 = 47

	// ISO5427Cyrillic is the uint16 identifier with IANA name ISO_5427.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO5427Cyrillic uint16 = 48

	// ISO42JISC62261978 is the uint16 identifier with IANA name JIS_C6226-1978.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO42JISC62261978 uint16 = 49

	// ISO47BSViewdata is the uint16 identifier with IANA name BS_viewdata.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO47BSViewdata uint16 = 50

	// ISO49INIS is the uint16 identifier with IANA name INIS.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO49INIS uint16 = 51

	// ISO50INIS8 is the uint16 identifier with IANA name INIS-8.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO50INIS8 uint16 = 52

	// ISO51INISCyrillic is the uint16 identifier with IANA name INIS-cyrillic.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO51INISCyrillic uint16 = 53

	// ISO54271981 is the uint16 identifier with IANA name ISO_5427:1981.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO54271981 uint16 = 54

	// ISO5428Greek is the uint16 identifier with IANA name ISO_5428:1980.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO5428Greek uint16 = 55

	// ISO57GB1988 is the uint16 identifier with IANA name GB_1988-80.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO57GB1988 uint16 = 56

	// ISO58GB231280 is the uint16 identifier with IANA name GB_2312-80.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO58GB231280 uint16 = 57

	// ISO61Norwegian2 is the uint16 identifier with IANA name NS_4551-2.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO61Norwegian2 uint16 = 58

	// ISO70VideotexSupp1 is the uint16 identifier with IANA name videotex-suppl.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO70VideotexSupp1 uint16 = 59

	// ISO84Portuguese2 is the uint16 identifier with IANA name PT2.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO84Portuguese2 uint16 = 60

	// ISO85Spanish2 is the uint16 identifier with IANA name ES2.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO85Spanish2 uint16 = 61

	// ISO86Hungarian is the uint16 identifier with IANA name MSZ_7795.3.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO86Hungarian uint16 = 62

	// ISO87JISX0208 is the uint16 identifier with IANA name JIS_C6226-1983.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO87JISX0208 uint16 = 63

	// ISO88Greek7 is the uint16 identifier with IANA name greek7.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO88Greek7 uint16 = 64

	// ISO89ASMO449 is the uint16 identifier with IANA name ASMO_449.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO89ASMO449 uint16 = 65

	// ISO90 is the uint16 identifier with IANA name iso-ir-90.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO90 uint16 = 66

	// ISO91JISC62291984a is the uint16 identifier with IANA name JIS_C6229-1984-a.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO91JISC62291984a uint16 = 67

	// ISO92JISC62991984b is the uint16 identifier with IANA name JIS_C6229-1984-b.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO92JISC62991984b uint16 = 68

	// ISO93JIS62291984badd is the uint16 identifier with IANA name JIS_C6229-1984-b-add.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO93JIS62291984badd uint16 = 69

	// ISO94JIS62291984hand is the uint16 identifier with IANA name JIS_C6229-1984-hand.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO94JIS62291984hand uint16 = 70

	// ISO95JIS62291984handadd is the uint16 identifier with IANA name JIS_C6229-1984-hand-add.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO95JIS62291984handadd uint16 = 71

	// ISO96JISC62291984kana is the uint16 identifier with IANA name JIS_C6229-1984-kana.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO96JISC62291984kana uint16 = 72

	// ISO2033 is the uint16 identifier with IANA name ISO_2033-1983.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO2033 uint16 = 73

	// ISO99NAPLPS is the uint16 identifier with IANA name ANSI_X3.110-1983.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO99NAPLPS uint16 = 74

	// ISO102T617bit is the uint16 identifier with IANA name T.61-7bit.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO102T617bit uint16 = 75

	// ISO103T618bit is the uint16 identifier with IANA name T.61-8bit.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO103T618bit uint16 = 76

	// ISO111ECMACyrillic is the uint16 identifier with IANA name ECMA-cyrillic.
	//
	// ISO registry
	// (formerly ECMA
	// registry )
	ISO111ECMACyrillic uint16 = 77

	// ISO121Canadian1 is the uint16 identifier with IANA name CSA_Z243.4-1985-1.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO121Canadian1 uint16 = 78

	// ISO122Canadian2 is the uint16 identifier with IANA name CSA_Z243.4-1985-2.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO122Canadian2 uint16 = 79

	// ISO123CSAZ24341985gr is the uint16 identifier with IANA name CSA_Z243.4-1985-gr.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO123CSAZ24341985gr uint16 = 80

	// ISO88596E is the uint16 identifier with IANA name ISO_8859-6-E (MIME: ISO-8859-6-E).
	//
	// rfc1556
	// Reference: RFC1556
	ISO88596E uint16 = 81

	// ISO88596I is the uint16 identifier with IANA name ISO_8859-6-I (MIME: ISO-8859-6-I).
	//
	// rfc1556
	// Reference: RFC1556
	ISO88596I uint16 = 82

	// ISO128T101G2 is the uint16 identifier with IANA name T.101-G2.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO128T101G2 uint16 = 83

	// ISO88598E is the uint16 identifier with IANA name ISO_8859-8-E (MIME: ISO-8859-8-E).
	//
	// rfc1556
	// Reference: RFC1556
	ISO88598E uint16 = 84

	// ISO88598I is the uint16 identifier with IANA name ISO_8859-8-I (MIME: ISO-8859-8-I).
	//
	// rfc1556
	// Reference: RFC1556
	ISO88598I uint16 = 85

	// ISO139CSN369103 is the uint16 identifier with IANA name CSN_369103.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO139CSN369103 uint16 = 86

	// ISO141JUSIB1002 is the uint16 identifier with IANA name JUS_I.B1.002.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO141JUSIB1002 uint16 = 87

	// ISO143IECP271 is the uint16 identifier with IANA name IEC_P27-1.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO143IECP271 uint16 = 88

	// ISO146Serbian is the uint16 identifier with IANA name JUS_I.B1.003-serb.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO146Serbian uint16 = 89

	// ISO147Macedonian is the uint16 identifier with IANA name JUS_I.B1.003-mac.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO147Macedonian uint16 = 90

	// ISO150GreekCCITT is the uint16 identifier with IANA name greek-ccitt.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO150GreekCCITT uint16 = 91

	// ISO151Cuba is the uint16 identifier with IANA name NC_NC00-10:81.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO151Cuba uint16 = 92

	// ISO6937Add is the uint16 identifier with IANA name ISO_6937-2-25.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO6937Add uint16 = 93

	// ISO153GOST1976874 is the uint16 identifier with IANA name GOST_19768-74.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO153GOST1976874 uint16 = 94

	// ISO8859Supp is the uint16 identifier with IANA name ISO_8859-supp.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO8859Supp uint16 = 95

	// ISO10367Box is the uint16 identifier with IANA name ISO_10367-box.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO10367Box uint16 = 96

	// ISO158Lap is the uint16 identifier with IANA name latin-lap.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO158Lap uint16 = 97

	// ISO159JISX02121990 is the uint16 identifier with IANA name JIS_X0212-1990.
	//
	// ISO-IR: International Register of Escape Sequences
	// Note: The current registration authority is IPSJ/ITSCJ, Japan.
	// Reference: RFC1345
	ISO159JISX02121990 uint16 = 98

	// ISO646Danish is the uint16 identifier with IANA name DS_2089.
	//
	// Danish Standard, DS 2089, February 1974
	// Reference: RFC1345
	ISO646Danish uint16 = 99

	// USDK is the uint16 identifier with IANA name us-dk.
	//
	// Reference: RFC1345
	USDK uint16 = 100

	// DKUS is the uint16 identifier with IANA name dk-us.
	//
	// Reference: RFC1345
	DKUS uint16 = 101

	// KSC5636 is the uint16 identifier with IANA name KSC5636.
	//
	// Reference: RFC1345
	KSC5636 uint16 = 102

	// Unicode11UTF7 is the uint16 identifier with IANA name UNICODE-1-1-UTF-7.
	//
	// rfc1642
	// Reference: RFC1642
	Unicode11UTF7 uint16 = 103

	// ISO2022CN is the uint16 identifier with IANA name ISO-2022-CN.
	//
	// rfc1922
	// Reference: RFC1922
	ISO2022CN uint16 = 104

	// ISO2022CNEXT is the uint16 identifier with IANA name ISO-2022-CN-EXT.
	//
	// rfc1922
	// Reference: RFC1922
	ISO2022CNEXT uint16 = 105

	// UTF8 is the uint16 identifier with IANA name UTF-8.
	//
	// rfc3629
	// Reference: RFC3629
	UTF8 uint16 = 106

	// ISO885913 is the uint16 identifier with IANA name ISO-8859-13.
	//
	// ISO See http://www.iana.org/assignments/charset-reg/ISO-8859-13 http://www.iana.org/assignments/charset-reg/ISO-8859-13
	ISO885913 uint16 = 109

	// ISO885914 is the uint16 identifier with IANA name ISO-8859-14.
	//
	// ISO See http://www.iana.org/assignments/charset-reg/ISO-8859-14
	ISO885914 uint16 = 110

	// ISO885915 is the uint16 identifier with IANA name ISO-8859-15.
	//
	// ISO
	// Please see: http://www.iana.org/assignments/charset-reg/ISO-8859-15
	ISO885915 uint16 = 111

	// ISO885916 is the uint16 identifier with IANA name ISO-8859-16.
	//
	// ISO
	ISO885916 uint16 = 112

	// GBK is the uint16 identifier with IANA name GBK.
	//
	// Chinese IT Standardization Technical Committee
	// Please see: http://www.iana.org/assignments/charset-reg/GBK
	GBK uint16 = 113

	// GB18030 is the uint16 identifier with IANA name GB18030.
	//
	// Chinese IT Standardization Technical Committee
	// Please see: http://www.iana.org/assignments/charset-reg/GB18030
	GB18030 uint16 = 114

	// OSDEBCDICDF0415 is the uint16 identifier with IANA name OSD_EBCDIC_DF04_15.
	//
	// Fujitsu-Siemens standard mainframe EBCDIC encoding
	// Please see: http://www.iana.org/assignments/charset-reg/OSD-EBCDIC-DF04-15
	OSDEBCDICDF0415 uint16 = 115

	// OSDEBCDICDF03IRV is the uint16 identifier with IANA name OSD_EBCDIC_DF03_IRV.
	//
	// Fujitsu-Siemens standard mainframe EBCDIC encoding
	// Please see: http://www.iana.org/assignments/charset-reg/OSD-EBCDIC-DF03-IRV
	OSDEBCDICDF03IRV uint16 = 116

	// OSDEBCDICDF041 is the uint16 identifier with IANA name OSD_EBCDIC_DF04_1.
	//
	// Fujitsu-Siemens standard mainframe EBCDIC encoding
	// Please see: http://www.iana.org/assignments/charset-reg/OSD-EBCDIC-DF04-1
	OSDEBCDICDF041 uint16 = 117

	// ISO115481 is the uint16 identifier with IANA name ISO-11548-1.
	//
	// See http://www.iana.org/assignments/charset-reg/ISO-11548-1
	ISO115481 uint16 = 118

	// KZ1048 is the uint16 identifier with IANA name KZ-1048.
	//
	// See http://www.iana.org/assignments/charset-reg/KZ-1048
	KZ1048 uint16 = 119

	// Unicode is the uint16 identifier with IANA name ISO-10646-UCS-2.
	//
	// the 2-octet Basic Multilingual Plane, aka Unicode
	// this needs to specify network byte order: the standard
	// does not specify (it is a 16-bit integer space)
	Unicode uint16 = 1000

	// UCS4 is the uint16 identifier with IANA name ISO-10646-UCS-4.
	//
	// the full code space. (same comment about byte order,
	// these are 31-bit numbers.
	UCS4 uint16 = 1001

	// UnicodeASCII is the uint16 identifier with IANA name ISO-10646-UCS-Basic.
	//
	// ASCII subset of Unicode.  Basic Latin = collection 1
	// See ISO 10646, Appendix A
	UnicodeASCII uint16 = 1002

	// UnicodeLatin1 is the uint16 identifier with IANA name ISO-10646-Unicode-Latin1.
	//
	// ISO Latin-1 subset of Unicode. Basic Latin and Latin-1
	// Supplement  = collections 1 and 2.  See ISO 10646,
	// Appendix A.  See rfc1815 .
	UnicodeLatin1 uint16 = 1003

	// UnicodeJapanese is the uint16 identifier with IANA name ISO-10646-J-1.
	//
	// ISO 10646 Japanese, see rfc1815 .
	UnicodeJapanese uint16 = 1004

	// UnicodeIBM1261 is the uint16 identifier with IANA name ISO-Unicode-IBM-1261.
	//
	// IBM Latin-2, -3, -5, Extended Presentation Set, GCSGID: 1261
	UnicodeIBM1261 uint16 = 1005

	// UnicodeIBM1268 is the uint16 identifier with IANA name ISO-Unicode-IBM-1268.
	//
	// IBM Latin-4 Extended Presentation Set, GCSGID: 1268
	UnicodeIBM1268 uint16 = 1006

	// UnicodeIBM1276 is the uint16 identifier with IANA name ISO-Unicode-IBM-1276.
	//
	// IBM Cyrillic Greek Extended Presentation Set, GCSGID: 1276
	UnicodeIBM1276 uint16 = 1007

	// UnicodeIBM1264 is the uint16 identifier with IANA name ISO-Unicode-IBM-1264.
	//
	// IBM Arabic Presentation Set, GCSGID: 1264
	UnicodeIBM1264 uint16 = 1008

	// UnicodeIBM1265 is the uint16 identifier with IANA name ISO-Unicode-IBM-1265.
	//
	// IBM Hebrew Presentation Set, GCSGID: 1265
	UnicodeIBM1265 uint16 = 1009

	// Unicode11 is the uint16 identifier with IANA name UNICODE-1-1.
	//
	// rfc1641
	// Reference: RFC1641
	Unicode11 uint16 = 1010

	// SCSU is the uint16 identifier with IANA name SCSU.
	//
	// SCSU See http://www.iana.org/assignments/charset-reg/SCSU
	SCSU uint16 = 1011

	// UTF7 is the uint16 identifier with IANA name UTF-7.
	//
	// rfc2152
	// Reference: RFC2152
	UTF7 uint16 = 1012

	// UTF16BE is the uint16 identifier with IANA name UTF-16BE.
	//
	// rfc2781
	// Reference: RFC2781
	UTF16BE uint16 = 1013

	// UTF16LE is the uint16 identifier with IANA name UTF-16LE.
	//
	// rfc2781
	// Reference: RFC2781
	UTF16LE uint16 = 1014

	// UTF16 is the uint16 identifier with IANA name UTF-16.
	//
	// rfc2781
	// Reference: RFC2781
	UTF16 uint16 = 1015

	// CESU8 is the uint16 identifier with IANA name CESU-8.
	//
	// https://www.unicode.org/unicode/reports/tr26
	CESU8 uint16 = 1016

	// UTF32 is the uint16 identifier with IANA name UTF-32.
	//
	// https://www.unicode.org/unicode/reports/tr19/
	UTF32 uint16 = 1017

	// UTF32BE is the uint16 identifier with IANA name UTF-32BE.
	//
	// https://www.unicode.org/unicode/reports/tr19/
	UTF32BE uint16 = 1018

	// UTF32LE is the uint16 identifier with IANA name UTF-32LE.
	//
	// https://www.unicode.org/unicode/reports/tr19/
	UTF32LE uint16 = 1019

	// BOCU1 is the uint16 identifier with IANA name BOCU-1.
	//
	// https://www.unicode.org/notes/tn6/
	BOCU1 uint16 = 1020

	// Windows30Latin1 is the uint16 identifier with IANA name ISO-8859-1-Windows-3.0-Latin-1.
	//
	// Extended ISO 8859-1 Latin-1 for Windows 3.0.
	// PCL Symbol Set id: 9U
	Windows30Latin1 uint16 = 2000

	// Windows31Latin1 is the uint16 identifier with IANA name ISO-8859-1-Windows-3.1-Latin-1.
	//
	// Extended ISO 8859-1 Latin-1 for Windows 3.1.
	// PCL Symbol Set id: 19U
	Windows31Latin1 uint16 = 2001

	// Windows31Latin2 is the uint16 identifier with IANA name ISO-8859-2-Windows-Latin-2.
	//
	// Extended ISO 8859-2.  Latin-2 for Windows 3.1.
	// PCL Symbol Set id: 9E
	Windows31Latin2 uint16 = 2002

	// Windows31Latin5 is the uint16 identifier with IANA name ISO-8859-9-Windows-Latin-5.
	//
	// Extended ISO 8859-9.  Latin-5 for Windows 3.1
	// PCL Symbol Set id: 5T
	Windows31Latin5 uint16 = 2003

	// HPRoman8 is the uint16 identifier with IANA name hp-roman8.
	//
	// LaserJet IIP Printer User's Manual,
	// HP part no 33471-90901, Hewlet-Packard, June 1989.
	// Reference: RFC1345
	HPRoman8 uint16 = 2004

	// AdobeStandardEncoding is the uint16 identifier with IANA name Adobe-Standard-Encoding.
	//
	// PostScript Language Reference Manual
	// PCL Symbol Set id: 10J
	AdobeStandardEncoding uint16 = 2005

	// VenturaUS is the uint16 identifier with IANA name Ventura-US.
	//
	// Ventura US.  ASCII plus characters typically used in
	// publishing, like pilcrow, copyright, registered, trade mark,
	// section, dagger, and double dagger in the range A0 (hex)
	// to FF (hex).
	// PCL Symbol Set id: 14J
	VenturaUS uint16 = 2006

	// VenturaInternational is the uint16 identifier with IANA name Ventura-International.
	//
	// Ventura International.  ASCII plus coded characters similar
	// to Roman8.
	// PCL Symbol Set id: 13J
	VenturaInternational uint16 = 2007

	// DECMCS is the uint16 identifier with IANA name DEC-MCS.
	//
	// VAX/VMS User's Manual,
	// Order Number: AI-Y517A-TE, April 1986.
	// Reference: RFC1345
	DECMCS uint16 = 2008

	// PC850Multilingual is the uint16 identifier with IANA name IBM850.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	PC850Multilingual uint16 = 2009

	// PC8DanishNorwegian is the uint16 identifier with IANA name PC8-Danish-Norwegian.
	//
	// PC Danish Norwegian
	// 8-bit PC set for Danish Norwegian
	// PCL Symbol Set id: 11U
	PC8DanishNorwegian uint16 = 2012

	// PC862LatinHebrew is the uint16 identifier with IANA name IBM862.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	PC862LatinHebrew uint16 = 2013

	// PC8Turkish is the uint16 identifier with IANA name PC8-Turkish.
	//
	// PC Latin Turkish.  PCL Symbol Set id: 9T
	PC8Turkish uint16 = 2014

	// IBMSymbols is the uint16 identifier with IANA name IBM-Symbols.
	//
	// Presentation Set, CPGID: 259
	IBMSymbols uint16 = 2015

	// IBMThai is the uint16 identifier with IANA name IBM-Thai.
	//
	// Presentation Set, CPGID: 838
	IBMThai uint16 = 2016

	// HPLegal is the uint16 identifier with IANA name HP-Legal.
	//
	// PCL 5 Comparison Guide, Hewlett-Packard,
	// HP part number 5961-0510, October 1992
	// PCL Symbol Set id: 1U
	HPLegal uint16 = 2017

	// HPPiFont is the uint16 identifier with IANA name HP-Pi-font.
	//
	// PCL 5 Comparison Guide, Hewlett-Packard,
	// HP part number 5961-0510, October 1992
	// PCL Symbol Set id: 15U
	HPPiFont uint16 = 2018

	// HPMath8 is the uint16 identifier with IANA name HP-Math8.
	//
	// PCL 5 Comparison Guide, Hewlett-Packard,
	// HP part number 5961-0510, October 1992
	// PCL Symbol Set id: 8M
	HPMath8 uint16 = 2019

	// HPPSMath is the uint16 identifier with IANA name Adobe-Symbol-Encoding.
	//
	// PostScript Language Reference Manual
	// PCL Symbol Set id: 5M
	HPPSMath uint16 = 2020

	// HPDesktop is the uint16 identifier with IANA name HP-DeskTop.
	//
	// PCL 5 Comparison Guide, Hewlett-Packard,
	// HP part number 5961-0510, October 1992
	// PCL Symbol Set id: 7J
	HPDesktop uint16 = 2021

	// VenturaMath is the uint16 identifier with IANA name Ventura-Math.
	//
	// PCL 5 Comparison Guide, Hewlett-Packard,
	// HP part number 5961-0510, October 1992
	// PCL Symbol Set id: 6M
	VenturaMath uint16 = 2022

	// MicrosoftPublishing is the uint16 identifier with IANA name Microsoft-Publishing.
	//
	// PCL 5 Comparison Guide, Hewlett-Packard,
	// HP part number 5961-0510, October 1992
	// PCL Symbol Set id: 6J
	MicrosoftPublishing uint16 = 2023

	// Windows31J is the uint16 identifier with IANA name Windows-31J.
	//
	// Windows Japanese.  A further extension of Shift_JIS
	// to include NEC special characters (Row 13), NEC
	// selection of IBM extensions (Rows 89 to 92), and IBM
	// extensions (Rows 115 to 119).  The CCS's are
	// JIS X0201:1997, JIS X0208:1997, and these extensions.
	// This charset can be used for the top-level media type "text",
	// but it is of limited or specialized use (see rfc2278 ).
	// PCL Symbol Set id: 19K
	Windows31J uint16 = 2024

	// GB2312 is the uint16 identifier with IANA name GB2312 (MIME: GB2312).
	//
	// Chinese for People's Republic of China (PRC) mixed one byte,
	// two byte set:
	// 20-7E = one byte ASCII
	// A1-FE = two byte PRC Kanji
	// See GB 2312-80
	// PCL Symbol Set Id: 18C
	GB2312 uint16 = 2025

	// Big5 is the uint16 identifier with IANA name Big5 (MIME: Big5).
	//
	// Chinese for Taiwan Multi-byte set.
	// PCL Symbol Set Id: 18T
	Big5 uint16 = 2026

	// Macintosh is the uint16 identifier with IANA name macintosh.
	//
	// The Unicode Standard ver1.0, ISBN 0-201-56788-1, Oct 1991
	// Reference: RFC1345
	Macintosh uint16 = 2027

	// IBM037 is the uint16 identifier with IANA name IBM037.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM037 uint16 = 2028

	// IBM038 is the uint16 identifier with IANA name IBM038.
	//
	// IBM 3174 Character Set Ref, GA27-3831-02, March 1990
	// Reference: RFC1345
	IBM038 uint16 = 2029

	// IBM273 is the uint16 identifier with IANA name IBM273.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM273 uint16 = 2030

	// IBM274 is the uint16 identifier with IANA name IBM274.
	//
	// IBM 3174 Character Set Ref, GA27-3831-02, March 1990
	// Reference: RFC1345
	IBM274 uint16 = 2031

	// IBM275 is the uint16 identifier with IANA name IBM275.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM275 uint16 = 2032

	// IBM277 is the uint16 identifier with IANA name IBM277.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM277 uint16 = 2033

	// IBM278 is the uint16 identifier with IANA name IBM278.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM278 uint16 = 2034

	// IBM280 is the uint16 identifier with IANA name IBM280.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM280 uint16 = 2035

	// IBM281 is the uint16 identifier with IANA name IBM281.
	//
	// IBM 3174 Character Set Ref, GA27-3831-02, March 1990
	// Reference: RFC1345
	IBM281 uint16 = 2036

	// IBM284 is the uint16 identifier with IANA name IBM284.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM284 uint16 = 2037

	// IBM285 is the uint16 identifier with IANA name IBM285.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM285 uint16 = 2038

	// IBM290 is the uint16 identifier with IANA name IBM290.
	//
	// IBM 3174 Character Set Ref, GA27-3831-02, March 1990
	// Reference: RFC1345
	IBM290 uint16 = 2039

	// IBM297 is the uint16 identifier with IANA name IBM297.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM297 uint16 = 2040

	// IBM420 is the uint16 identifier with IANA name IBM420.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990,
	// IBM NLS RM p 11-11
	// Reference: RFC1345
	IBM420 uint16 = 2041

	// IBM423 is the uint16 identifier with IANA name IBM423.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM423 uint16 = 2042

	// IBM424 is the uint16 identifier with IANA name IBM424.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM424 uint16 = 2043

	// PC8CodePage437 is the uint16 identifier with IANA name IBM437.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	PC8CodePage437 uint16 = 2011

	// IBM500 is the uint16 identifier with IANA name IBM500.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM500 uint16 = 2044

	// IBM851 is the uint16 identifier with IANA name IBM851.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM851 uint16 = 2045

	// PCp852 is the uint16 identifier with IANA name IBM852.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	PCp852 uint16 = 2010

	// IBM855 is the uint16 identifier with IANA name IBM855.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM855 uint16 = 2046

	// IBM857 is the uint16 identifier with IANA name IBM857.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM857 uint16 = 2047

	// IBM860 is the uint16 identifier with IANA name IBM860.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM860 uint16 = 2048

	// IBM861 is the uint16 identifier with IANA name IBM861.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM861 uint16 = 2049

	// IBM863 is the uint16 identifier with IANA name IBM863.
	//
	// IBM Keyboard layouts and code pages, PN 07G4586 June 1991
	// Reference: RFC1345
	IBM863 uint16 = 2050

	// IBM864 is the uint16 identifier with IANA name IBM864.
	//
	// IBM Keyboard layouts and code pages, PN 07G4586 June 1991
	// Reference: RFC1345
	IBM864 uint16 = 2051

	// IBM865 is the uint16 identifier with IANA name IBM865.
	//
	// IBM DOS 3.3 Ref (Abridged), 94X9575 (Feb 1987)
	// Reference: RFC1345
	IBM865 uint16 = 2052

	// IBM868 is the uint16 identifier with IANA name IBM868.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM868 uint16 = 2053

	// IBM869 is the uint16 identifier with IANA name IBM869.
	//
	// IBM Keyboard layouts and code pages, PN 07G4586 June 1991
	// Reference: RFC1345
	IBM869 uint16 = 2054

	// IBM870 is the uint16 identifier with IANA name IBM870.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM870 uint16 = 2055

	// IBM871 is the uint16 identifier with IANA name IBM871.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM871 uint16 = 2056

	// IBM880 is the uint16 identifier with IANA name IBM880.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM880 uint16 = 2057

	// IBM891 is the uint16 identifier with IANA name IBM891.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM891 uint16 = 2058

	// IBM903 is the uint16 identifier with IANA name IBM903.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM903 uint16 = 2059

	// IBBM904 is the uint16 identifier with IANA name IBM904.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBBM904 uint16 = 2060

	// IBM905 is the uint16 identifier with IANA name IBM905.
	//
	// IBM 3174 Character Set Ref, GA27-3831-02, March 1990
	// Reference: RFC1345
	IBM905 uint16 = 2061

	// IBM918 is the uint16 identifier with IANA name IBM918.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM918 uint16 = 2062

	// IBM1026 is the uint16 identifier with IANA name IBM1026.
	//
	// IBM NLS RM Vol2 SE09-8002-01, March 1990
	// Reference: RFC1345
	IBM1026 uint16 = 2063

	// IBMEBCDICATDE is the uint16 identifier with IANA name EBCDIC-AT-DE.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	IBMEBCDICATDE uint16 = 2064

	// EBCDICATDEA is the uint16 identifier with IANA name EBCDIC-AT-DE-A.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICATDEA uint16 = 2065

	// EBCDICCAFR is the uint16 identifier with IANA name EBCDIC-CA-FR.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICCAFR uint16 = 2066

	// EBCDICDKNO is the uint16 identifier with IANA name EBCDIC-DK-NO.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICDKNO uint16 = 2067

	// EBCDICDKNOA is the uint16 identifier with IANA name EBCDIC-DK-NO-A.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICDKNOA uint16 = 2068

	// EBCDICFISE is the uint16 identifier with IANA name EBCDIC-FI-SE.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICFISE uint16 = 2069

	// EBCDICFISEA is the uint16 identifier with IANA name EBCDIC-FI-SE-A.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICFISEA uint16 = 2070

	// EBCDICFR is the uint16 identifier with IANA name EBCDIC-FR.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICFR uint16 = 2071

	// EBCDICIT is the uint16 identifier with IANA name EBCDIC-IT.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICIT uint16 = 2072

	// EBCDICPT is the uint16 identifier with IANA name EBCDIC-PT.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICPT uint16 = 2073

	// EBCDICES is the uint16 identifier with IANA name EBCDIC-ES.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICES uint16 = 2074

	// EBCDICESA is the uint16 identifier with IANA name EBCDIC-ES-A.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICESA uint16 = 2075

	// EBCDICESS is the uint16 identifier with IANA name EBCDIC-ES-S.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICESS uint16 = 2076

	// EBCDICUK is the uint16 identifier with IANA name EBCDIC-UK.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICUK uint16 = 2077

	// EBCDICUS is the uint16 identifier with IANA name EBCDIC-US.
	//
	// IBM 3270 Char Set Ref Ch 10, GA27-2837-9, April 1987
	// Reference: RFC1345
	EBCDICUS uint16 = 2078

	// Unknown8BiT is the uint16 identifier with IANA name UNKNOWN-8BIT.
	//
	// Reference: RFC1428
	Unknown8BiT uint16 = 2079

	// Mnemonic is the uint16 identifier with IANA name MNEMONIC.
	//
	// rfc1345 , also known as "mnemonic+ascii+38"
	// Reference: RFC1345
	Mnemonic uint16 = 2080

	// Mnem is the uint16 identifier with IANA name MNEM.
	//
	// rfc1345 , also known as "mnemonic+ascii+8200"
	// Reference: RFC1345
	Mnem uint16 = 2081

	// VISCII is the uint16 identifier with IANA name VISCII.
	//
	// rfc1456
	// Reference: RFC1456
	VISCII uint16 = 2082

	// VIQR is the uint16 identifier with IANA name VIQR.
	//
	// rfc1456
	// Reference: RFC1456
	VIQR uint16 = 2083

	// KOI8R is the uint16 identifier with IANA name KOI8-R (MIME: KOI8-R).
	//
	// rfc1489 , based on GOST-19768-74, ISO-6937/8,
	// INIS-Cyrillic, ISO-5427.
	// Reference: RFC1489
	KOI8R uint16 = 2084

	// HZGB2312 is the uint16 identifier with IANA name HZ-GB-2312.
	//
	// rfc1842 , rfc1843 rfc1843 rfc1842
	HZGB2312 uint16 = 2085

	// IBM866 is the uint16 identifier with IANA name IBM866.
	//
	// IBM NLDG Volume 2 (SE09-8002-03) August 1994
	IBM866 uint16 = 2086

	// PC775Baltic is the uint16 identifier with IANA name IBM775.
	//
	// HP PCL 5 Comparison Guide (P/N 5021-0329) pp B-13, 1996
	PC775Baltic uint16 = 2087

	// KOI8U is the uint16 identifier with IANA name KOI8-U.
	//
	// rfc2319
	// Reference: RFC2319
	KOI8U uint16 = 2088

	// IBM00858 is the uint16 identifier with IANA name IBM00858.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM00858
	IBM00858 uint16 = 2089

	// IBM00924 is the uint16 identifier with IANA name IBM00924.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM00924
	IBM00924 uint16 = 2090

	// IBM01140 is the uint16 identifier with IANA name IBM01140.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01140
	IBM01140 uint16 = 2091

	// IBM01141 is the uint16 identifier with IANA name IBM01141.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01141
	IBM01141 uint16 = 2092

	// IBM01142 is the uint16 identifier with IANA name IBM01142.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01142
	IBM01142 uint16 = 2093

	// IBM01143 is the uint16 identifier with IANA name IBM01143.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01143
	IBM01143 uint16 = 2094

	// IBM01144 is the uint16 identifier with IANA name IBM01144.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01144
	IBM01144 uint16 = 2095

	// IBM01145 is the uint16 identifier with IANA name IBM01145.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01145
	IBM01145 uint16 = 2096

	// IBM01146 is the uint16 identifier with IANA name IBM01146.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01146
	IBM01146 uint16 = 2097

	// IBM01147 is the uint16 identifier with IANA name IBM01147.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01147
	IBM01147 uint16 = 2098

	// IBM01148 is the uint16 identifier with IANA name IBM01148.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01148
	IBM01148 uint16 = 2099

	// IBM01149 is the uint16 identifier with IANA name IBM01149.
	//
	// IBM See http://www.iana.org/assignments/charset-reg/IBM01149
	IBM01149 uint16 = 2100

	// Big5HKSCS is the uint16 identifier with IANA name Big5-HKSCS.
	//
	// See http://www.iana.org/assignments/charset-reg/Big5-HKSCS
	Big5HKSCS uint16 = 2101

	// IBM1047 is the uint16 identifier with IANA name IBM1047.
	//
	// IBM1047 (EBCDIC Latin 1/Open Systems) http://www-1.ibm.com/servers/eserver/iseries/software/globalization/pdf/cp01047z.pdf
	IBM1047 uint16 = 2102

	// PTCP154 is the uint16 identifier with IANA name PTCP154.
	//
	// See http://www.iana.org/assignments/charset-reg/PTCP154
	PTCP154 uint16 = 2103

	// Amiga1251 is the uint16 identifier with IANA name Amiga-1251.
	//
	// See http://www.amiga.ultranet.ru/Amiga-1251.html
	Amiga1251 uint16 = 2104

	// KOI7switched is the uint16 identifier with IANA name KOI7-switched.
	//
	// See http://www.iana.org/assignments/charset-reg/KOI7-switched
	KOI7switched uint16 = 2105

	// BRF is the uint16 identifier with IANA name BRF.
	//
	// See http://www.iana.org/assignments/charset-reg/BRF
	BRF uint16 = 2106

	// TSCII is the uint16 identifier with IANA name TSCII.
	//
	// See http://www.iana.org/assignments/charset-reg/TSCII
	TSCII uint16 = 2107

	// CP51932 is the uint16 identifier with IANA name CP51932.
	//
	// See http://www.iana.org/assignments/charset-reg/CP51932
	CP51932 uint16 = 2108

	// Windows874 is the uint16 identifier with IANA name windows-874.
	//
	// See http://www.iana.org/assignments/charset-reg/windows-874
	Windows874 uint16 = 2109

	// Windows1250 is the uint16 identifier with IANA name windows-1250.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1250
	Windows1250 uint16 = 2250

	// Windows1251 is the uint16 identifier with IANA name windows-1251.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1251
	Windows1251 uint16 = 2251

	// Windows1252 is the uint16 identifier with IANA name windows-1252.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1252
	Windows1252 uint16 = 2252

	// Windows1253 is the uint16 identifier with IANA name windows-1253.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1253
	Windows1253 uint16 = 2253

	// Windows1254 is the uint16 identifier with IANA name windows-1254.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1254
	Windows1254 uint16 = 2254

	// Windows1255 is the uint16 identifier with IANA name windows-1255.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1255
	Windows1255 uint16 = 2255

	// Windows1256 is the uint16 identifier with IANA name windows-1256.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1256
	Windows1256 uint16 = 2256

	// Windows1257 is the uint16 identifier with IANA name windows-1257.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1257
	Windows1257 uint16 = 2257

	// Windows1258 is the uint16 identifier with IANA name windows-1258.
	//
	// Microsoft http://www.iana.org/assignments/charset-reg/windows-1258
	Windows1258 uint16 = 2258

	// TIS620 is the uint16 identifier with IANA name TIS-620.
	//
	// Thai Industrial Standards Institute (TISI)
	TIS620 uint16 = 2259

	// CP50220 is the uint16 identifier with IANA name CP50220.
	//
	// See http://www.iana.org/assignments/charset-reg/CP50220
	CP50220 uint16 = 2260
)

//Name - string of code page name
var Name = map[uint16]string{
	ASCII:            "ASCII",
	ISOLatin1:        "ISOLatin1",
	ISOLatin2:        "ISOLatin2",
	ISOLatin3:        "ISOLatin3",
	ISOLatin4:        "ISOLatin4",
	ISOLatinCyrillic: "ISOLatinCyrillic",
	ISOLatinArabic:   "ISOLatinArabic",
	ISOLatinGreek:    "ISOLatinGreek",
	ISOLatinHebrew:   "ISOLatinHebrew",
	ISOLatin5:        "ISOLatin5",
	ISOLatin6:        "ISOLatin6",
	IBM866:           "IBM866",
	Windows1251:      "Windows1251",
	UTF8:             "UTF8",
	UTF16:            "UTF16",
	UTF32:            "UTF32",
	KOI8R:            "KOI8R",
}
