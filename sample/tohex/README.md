# sample to package cpd #

(c) softlandia@gmail.com

tohex -- simple program using cpd to encode the input string to the specified encoding and return the string from the hexadecimal code of the received runes

sample:
>tohex Hi utf-16le
>"\x48\x0\x69\x0"

>tohex рс-Иџ utf16le
>"\x30\x4\x31\x4\x2D\x0\x51\x4\x4F\x4"

>tohex.exe рс-Иџ win1251 >r.txt
>"\xE0\xE1\x2D\xB8\xFF"

result string usefull using in golang code

