package parser

import (
	"errors"
	"sort"
)

func checkPupu(pupu []int, max int) (e error) {
	var tmp []int

	tmp = make([]int, len(pupu))
	copy(tmp, pupu)
	sort.Ints(tmp)
	for i := 0; i < len(pupu)-1; i++ {
		if tmp[i] < 0 {
			return errors.New("please don't use negative number")
		}
		if tmp[i] == tmp[i+1] {
			return errors.New("duplicate number")
		}
	}
	if tmp[len(tmp)-1] >= max {
		return errors.New("bad number: exceeds max")
	}
	if len(tmp) != max {
		return errors.New("parsing error: incorrect number of lines")
	}
	return
}
