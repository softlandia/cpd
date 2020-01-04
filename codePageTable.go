package cpd

//codePageTable

// return index of rune in code page table
// return 0 if rune not in code page table
func (t *codePageTable) index(r rune) int {
	for j, e := range *t {
		if r == e.code {
			return j
		}
	}
	return 0
}

func (t *codePageTable) clear() {
	for i := 0; i < len(t); i++ {
		t[i].count = 0
	}
}

// founded - calculates total number of matching
func (t *codePageTable) founded() (res int) {
	//0 элемент исключён, он не содержит количество найденных букв
	for i := 1; i < len(t); i++ {
		res += t[i].count
	}
	return
}
