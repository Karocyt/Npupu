package parser

import "sort"

func check_pupu(pupu []int, max int) bool{
	var tmp []int
	tmp = make([]int, len(pupu))
	copy(tmp, pupu)
	sort.Ints(tmp)
	for i := 0; i < len(pupu) - 1; i++ {
		if tmp[i] == tmp[i + 1]{
			return false
		}
	}
	if tmp[len(tmp) - 1] >= max { return false }
	return true
}