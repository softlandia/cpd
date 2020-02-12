package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/softlandia/cpd"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("using: 'tohex data CPXXX'\n")
		os.Exit(0)
	}
	//Проверим что целевая кодировка поддерживается
	if !cpd.SupportedEncoder(os.Args[2]) {
		fmt.Printf("landing codepage: '%s' not supported\n", os.Args[2])
		os.Exit(0)
	}
	//создаём новый ридер, при чтении из него будем получать уже в новой кодровке
	r, err := cpd.NewReaderTo(strings.NewReader(os.Args[1]), os.Args[2])
	if err != nil {
		fmt.Printf("cpd.NewReaderTo return error: '%v' \n", err)
		os.Exit(0)
	}
	//читаем данные
	b, _ := ioutil.ReadAll(r)
	var sb strings.Builder
	//собираем строку содержащую шестнадцатеричные коды символов в целевой кодировке
	sb.WriteString("\"")
	for _, r := range b {
		sb.WriteString("\\x")
		sb.WriteString(fmt.Sprintf("%X", r))
	}
	sb.WriteString("\"")
	fmt.Printf("%s", sb.String())
}
