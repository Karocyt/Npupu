package parser

import (
	"github.com/Karocyt/Npupu/internal/heuristics"
	"fmt"

)

func countInv(pupu []int, size int) (invCount int, posEmpty int) {
	_, goal := heuristics.MakeGoal(size)

	get1D := func (lol int) int {

		x, y := goal[lol][0], goal[lol][1]
		fmt.Println(heuristics.Get1d(x, y, size))
		return  heuristics.Get1d(x, y, size)
	}
	posEmpty = size


	for i := 0; i < size * size - 1; i++ {
		for j := (i + 1); j < size*size; j++ {
			posN1 := get1D(pupu[i])
			fmt.Println(posN1, i, pupu[i])
			posN2 := get1D(pupu[j])
			fmt.Println(posN2, i, pupu[j])
			if pupu[j] != 0 && pupu[i] != 0 && posN1 > posN2 {
				invCount++
			}
			if pupu[j] == 0 && j != size * size - 1 {
				posEmpty = size - (j / size)
				posN1 = get1D(pupu[i])
				posN2 = get1D(pupu[j + 1])
				if posN1 > posN2 {
					invCount++
				}
			}
		}
	}
	return
}

func checkSolvy(pupu []int, size int) bool{
	invCount, posEmpty := countInv(pupu, size)
	fmt.Println("info", invCount, posEmpty)
	if size % 2 == 1 && invCount % 2 == 0 {
			return true
	} else if  size % 2 == 0 && (posEmpty % 2 != invCount % 2) {
		return true
	} else {
		return false
	}
}