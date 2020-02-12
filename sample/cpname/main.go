package main

import (
	"fmt"

	"github.com/softlandia/cpd"

	hi "golang.org/x/text/encoding/htmlindex"
)

func main() {
	fmt.Printf("'github.com/softlandia/cpd' => 'golang.org/x/text/encoding/htmlindex'\n")
	for cp := range cpd.NewCodepageDic() {
		e, err := hi.Get(cp.String())
		if err == nil {
			name, _ := hi.Name(e)
			fmt.Printf("%s => %s\n", cp, name)
		} else {
			fmt.Printf("%s not present in 'htmlindex'\n", cp)
		}
	}
}
