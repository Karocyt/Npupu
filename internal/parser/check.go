package parser

import (
	"errors"
	"sort"
)

func check_pupu(pupu []int, max int) (e error){
	var tmp []int

	tmp = make([]int, len(pupu))
	copy(tmp, pupu)
	sort.Ints(tmp)
	for i := 0; i < len(pupu) - 1; i++ {
		if tmp[i] == tmp[i + 1]{
			e = errors.New("duplicate number")
			return
		}
	}
	if tmp[len(tmp) - 1] >= max {
		e = errors.New("bad number: exceeds max")
		return
	}
	return
}