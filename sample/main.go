package main

import (
	"fmt"
	"os"

	"github.com/softlandia/cpd"
)

func main() {
	t, _ := cpd.FileCodePageDetect(os.Args[1])
	fmt.Printf("cpd.FileCodePageDetect():\t%s\n", t)
	for id, cp := range cpd.CodepageDic {
		fmt.Printf("%s\tmatches:%s\t%s\n", id, cp.MatchRes, cp.MatchingRunes())
	}
}
