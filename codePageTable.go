package cpd

//codePageTable

//return index of rune in code page table
//return 0 if rune not in code page table
//TODO надо переименовать, сейчас звучит как метод возвращающий true/false, должно быть типа index()
func (t *codePageTable) containsRune(r rune) int {
	for j, e := range *t {
		if r == e.code {
			return j
		}
	}
	return 0
}

func (t *codePageTable) isUpper(r rune) bool {
	for i := 10; i < len(t); i++ {
		if r == (*t)[i].code {
			return true
		}
	}
	return false
}
