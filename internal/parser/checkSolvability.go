package parser

import (
	"github.com/Karocyt/Npupu/internal/heuristics"
)

func goBotRight(pupu []int, goal map[int][2]int, size int) ([]int, map[int][2]int) {
	iz := -1
	for i, v := range pupu {
		if 0 == v {
			iz = i
		}
	}
	for iz/size+1 != size {
		tmp := pupu[iz]
		pupu[iz] = pupu[iz+size]
		pupu[iz+size] = tmp
		iz += size
	}
	for iz != size*size-1 {
		tmp := pupu[iz]
		pupu[iz] = pupu[iz+1]
		pupu[iz+1] = tmp
		iz++
	}
	for i, v := range pupu {
		x, y := heuristics.Get2d(i, size)
		goal[v] = [2]int{x, y}
	}
	return pupu, goal
}

func countInv(pupu []int, size int) (invCount int, distEmpty int) {
	pupu_f, goal := heuristics.MakeGoal(size)
	pupu_f, goal = goBotRight(pupu_f, goal, size)
	get1D := func(lol int) int {
		x, y := goal[lol][0], goal[lol][1]
		return heuristics.Get1d(x, y, size)
	}
	distEmpty = size

	for i := 0; i < size*size-1; i++ {
		for j := (i + 1); j < size*size; j++ {
			posN1 := get1D(pupu[i])
			posN2 := get1D(pupu[j])
			if pupu[i] != 0 && pupu[j] != 0 && posN1 > posN2 {
				invCount++
			}
			if pupu[j] == 0 && j != size*size-1 {
				distEmpty = size - (j / size)
			}
		}
	}

	return
}

func checkSolvy(pupu []int, size int) bool {
	invCount, distEmpty := countInv(pupu, size)

	if size%2 == 1 {
		return invCount%2 == 0
	} else {
		if distEmpty%2 == 0 {
			return invCount%2 == 1
		} else {
			return invCount%2 == 0
		}
	}
}
