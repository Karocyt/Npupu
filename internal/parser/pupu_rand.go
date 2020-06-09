package parser

import (
	"fmt"
	"github.com/Karocyt/Npupu/internal/heuristics"
	"math/rand"
	"time"
)

func pupu_rand() ([]int, int){
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(3) + 3
	nbMove := rand.Intn(150) + 50
	pupu, _ := heuristics.MakeGoal(size)
	var iz int
	for i, v := range pupu {
		if 0 == v {
			iz = i
		}
	}
	for nbMove > 0 {
		randDir := rand.Intn(4)

		if randDir == 0 {
			if iz/size+1 != size {
				tmp := pupu[iz]
				pupu[iz] = pupu[iz+size]
				pupu[iz+size] = tmp
				iz += size
			}
		}

		if randDir == 1 {
			if iz % size != size  - 1 {
				tmp := pupu[iz]
				pupu[iz] = pupu[iz+1]
				pupu[iz+1] = tmp
				iz++
			}
		}

		if randDir == 2 {
			if iz > size {
				tmp := pupu[iz]
				pupu[iz] = pupu[iz - size]
				pupu[iz - size] = tmp
				iz -= size
			}
		}

		if randDir == 3 {
			if iz % size != 0 {
				tmp := pupu[iz]
				pupu[iz] = pupu[iz-1]
				pupu[iz-1] = tmp
				iz--
			}
		}
		nbMove--
	}
	fmt.Printf("Auto generate pupu size: %d\n", size)
	return pupu, size
}