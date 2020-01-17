# code page detect #

(c) softlandia@gmail.com

>download: go get -u github.com/softlandia/cpd  
>install: go install

библиотека для golang

предназначена для автоматического определения кодовой страницы текстовых файлов или потоков байт  
поддерживает следующие кодовые страницы:

no ID                Name		uint16

1. ASCII:            "ASCII",		3
2. ISOLatinCyrillic: "ISO-8859-5",	8
3. CP866:            "CP866",		2086
4. Windows1251:      "Windows-1251",	2251
5. UTF8:             "UTF-8",		106
6. UTF16LE:          "UTF-16LE",	1014
7. UTF16BE:          "UTF-16BE",	1013
8. KOI8R:            "KOI8-R",		2084
9. UTF32LE:          "UTF-32LE",	1019
10. UTF32BE:         "UTF-32BE",	1018

## особенности ##

определение делается как по наличию признака BOM в начале файла так и по эвристическому алгоритму  
если данные содержат только латинские символы (первая половина ASCII таблицы) будет определена кодировка UTF-8  
это не является ошибкой, поскольку такой файл или данные действительно можно корректно интерпретировать как UTF-8

при использовании golang 1.12.6 в проект добавляется код размером ~250 kB

ВНИМАНИЕ!  
файлы без BOM в кодировке UTF16le и UTF16be при отсутсвии русских букв опознаются не верно

## зависимости ##

>"golang.org/x/text/encoding/charmap"  
>"golang.org/x/text/transform"

## типы ##

IDCodePage uint16 - индекс кодовой страницы, значения взяты из файла поставки golang golang.org\x\text\encoding\internal\identifier\mib.go
поддерживается interface String(), и можно выводить так
    cp := cpd.UTF8
    fmt.Printf("code page index, name: %d, %s\n", cp, cp)
    //>code page index, name: 106, UTF-8

## глобальные переменные ##

ReadBufSize int = 1024 // количество байт считываемых из ридера (буфера) для определения кодировки

## функции ##

1. CodePageDetect(r io.Reader, stopStr ...string) (IDCodePage, error)
2. FileCodePageDetect(fn string, stopStr ...string) (IDCodePage, error)

## описание ##

    func CodePageAutoDetect(content []byte) (result IDCodePage) 
      автоматическое определеие кодировки по входному слайсу байт
      использовать вместо golang.org/x/net/html/charset.DetermineEncoding()

    CodePageDetect(r io.Reader, stopStr ...string) (IDCodePage, error)
      определяет кодовую страницу считывая поток байтов из 'r' 
      используется 'reflect.ValueOf(r).IsValid()' для проверки 'r' на существование
      считывает из 'r' первые ReadBufSize байтов
      параметр stopStr пока не используется

    FileCodePageDetect(fn string, stopStr ...string) (IDCodePage, error)
      определяет кодовую страницу считывая файл 'fn', считывает из файла первые ReadBufSize байтов
      ошибку возвращает если проблемы с открытием файла 'fn'
      возвращает cpd.ASCII если колировка не определена

## tests & static analiz ##

coverage: 89% of statements  
в папке "test_files" лежат файлы для тестов, соответственно не править и не удалять

linter.md отчёт статического анализатора golangci-lint
