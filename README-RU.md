# code page detect #

(c) softlandia@gmail.com

>download: go get -u github.com/softlandia/cpd  
>install: go install

библиотека для golang

предназначена для автоматического определения кодовой страницы текстовых файлов или потоков байт  
поддерживает следующие кодовые страницы:

no ID                Name

1. ASCII:            "ASCII",
2. ISOLatinCyrillic: "ISO-8859-5",
3. CP866:            "CP866",
4. Windows1251:      "Windows-1251",
5. UTF8:             "UTF-8",
6. UTF16LE:          "UTF-16LE",
7. UTF16BE:          "UTF-16BE",
8. UTF32:            "UTF-32",
9. KOI8R:            "KOI8-R",
10. Unicode:          "Unicode",
11. UTF7:             "UTF-7",
12. UTF32LE:          "UTF-32LE",
13. UTF32BE:          "UTF-32BE",

## особенности ##

если данные содержат только латинские символы (первая половина ASCII таблицы) будет определена кодировка UTF-8  
это не является ошибкой, поскольку такой файл или данные действительно можно использовать как UTF-8

при использовании golang 1.12.6 в проект добавляется код размером ~240 kB

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

    CodePageDetect(r io.Reader, stopStr ...string) (IDCodePage, error)
      определяет кодовую страницу считывая поток байтов из 'r' 
      используется 'reflect.ValueOf(r).IsValid()' для проверки 'r' на существование
      считывает из 'r' первые ReadBufSize байтов
      параметр stopStr пока не используется

    FileCodePageDetect(fn string, stopStr ...string) (IDCodePage, error)
      определяет кодовую страницу считывая файл 'fn', считывает из файла первые ReadBufSize байтов
      ошибку возвращает если проблемы с открытием файла 'fn'
      возвращает cpd.ASCII если колировка не определена

## tests ##

coverage: 84.0% of statements  
в папке "test_files" лежат файлы для тестов, соответственно не править и не удалять
