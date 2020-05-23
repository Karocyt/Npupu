package solverv2

import (
	"fmt"
)

var count int
var nb_open int

func appendAllNewStates(solver *Solver,  depth int) *Solver {
	fmt.Println("in:" ,solver, solver.vu)
	if solver.vu == true {
		/*if len(solver.kid) == 0 {
			if solver.dady != nil {
				return *solver.dady
			}
		}*/
		for i := 0; i < len(solver.kid) ; i++ {
				if solver.kid[i].vu == false {
					fmt.Println("out" ,solver.kid[i], solver.kid[i].vu)
					return &solver.kid[i]
				}
		}
		if solver.dady != nil {
			fmt.Println("out dady" ,*solver.dady,  (*solver.dady).vu)
			tmp :=  *solver.dady
			return &tmp
		}
	}
	solver.vu = true

	state := solver.state
	key := state.mapKey()
	if explored[key] == false {
		explored[key] = true
		if solver.open == false {

			nextStates := make([]gridState, 0, 4)
			state = solver.state
			solver.depth = state.depth
			for _, newState := range state.generateNextStates() {
				if hasSeen(newState) == false {
					nb_open++
					fmt.Println("nb_open:", nb_open)
					newState.score = solver.fn(newState.grid, size)
					nextStates = append(nextStates, newState)
				}
			}

			solver.kid = make([]Solver, len(nextStates))
			for i := 0; i < len(nextStates); i++ {

				solver.kid[i].depth = depth + 1
				solver.kid[i].fn = solver.fn
				solver.kid[i].state = nextStates[i]
				solver.kid[i].dady = solver
				solver.kid[i].id = i
			}
			solver.open = true
		}
	}
	/*	fmt.Println(solver, "score: ", solver.state.score)
		if solver.vu == true {
			if solver.dady != nil {
				tmp := *solver.dady
				for i := 0; i < len(tmp.kid); i++ {
					if tmp.kid[i].vu == false {
						fmt.Println("up")
						return appendAllNewStates(tmp.kid[i], stop, deap, nb_open)
					}
				}
			}
		}
		if solver.state.score == 0 {
			stop = 1
			fmt.Println("win")
			return solver
		}
		if stop == 1{
			solver.state.score = -1
			return solver
		}
		if solver.dady != nil {
			if explored[solver.state.mapKey()] == true {
					fmt.Println("up")

					return appendAllNewStates(*solver.dady, stop, deap, nb_open)
				}
			}

		//fmt.Println("stop: ", stop, solver.depth)
		if stop == 1 {
			fmt.Println("stop: ", stop)
			if solver.state.score == 0 {
				fmt.Println("win")
				return solver
			} else {
				stop = 0
			}
		}
		//fmt.Println("lol1")

	//	fmt.Println(explored)


		//for len(nextStates) == 0 && len(solver.openedStates) > 0 {
		//	fmt.Println("kéblo", solver.kid)



		//}

		if solver.dady != nil {
			if len(solver.kid) == 0 {
				fmt.Println("up")
				return appendAllNewStates(*solver.dady, stop, deap, nb_open)
			}
		}
		if len(nextStates) == 0 {
			return solver
		}

	//fmt.Println(solver)
		for i := 0; i < len(solver.kid); i++ {
		//	fmt.Println("lol", explored[solver.kid[i].openedStates[0].mapKey()])
			if explored[solver.kid[i].state.mapKey()] == true {
				continue
			}
			if explored[solver.kid[i].state.mapKey()] == false {
				if solver.vu == true {
					continue
				}
				if solver.kid[i].state.score == 0 {
					return solver.kid[i]
				}
				fmt.Println("kid: ", solver.kid[i])
				solver.vu = true
				return appendAllNewStates(solver.kid[i], stop, deap, nb_open)
				fmt.Println("lol   LOL   lol")
				if explored[solver.kid[i].openedStates[0].mapKey()] == false {
					if solver.kid[i].openedStates[0].score == 0 {
						stop = 1
						return solver
					}
						fmt.Println(solver.kid[i])

					appendAllNewStates(solver.kid[i], stop)
				}
			}
		}
		if solver.dady != nil {
			fmt.Println("up")
			return appendAllNewStates(*solver.dady, stop, deap, nb_open)
		}
			if len(nextStates) == 0 && len(solver.openedStates) > 0 {
				solver.openedStates = solver.openedStates[0 : len(solver.openedStates)-1]
				if len(solver.Solution) > 0 {
					solver.Solution = solver.Solution[0 : len(solver.Solution)-1]
				}
			}
		}

		for i := 0; i < len(nextStates); i++ {
			if nextStates[i].score == 0 {
				stop = 1
			}
			 if solver.explored[nextStates[i].mapKey()] == false {

			solver.Solution = append(solver.Solution, nextStates[i])
				solver.openedStates = append(solver.openedStates, nextStates[i])
				solver.totalOpenedStates++
				solver.depth = nextStates[i].depth
				solver.appendAllNewStates()
			}



		sort.Slice(nextStates, func(i, j int) bool {
			return nextStates[i].score > nextStates[j].score
		})
		//fmt.Println(state)
		for _, newState := range nextStates {
			//fmt.Println(newState.score)
			solver.Solution = append(solver.Solution, newState)
			solver.openedStates = append(solver.openedStates, newState)
			solver.totalOpenedStates++
			solver.depth = newState.depth
		}
		//fmt.Println()
		if len(solver.openedStates) > solver.maxOpenedStates {
			solver.maxOpenedStates = len(solver.openedStates)
		}



		if state.score != 0 && solver.explored[nextStates[0].mapKey()] == false && stop == 0{
			if state.score == 0 {
				stop = 1
			}
			solver.appendAllNewStates()
		} */

	if solver.dady != nil {
		tmp := *solver.dady
		for i := 0; i < len(tmp.kid) ; i++ {
			if tmp.kid[i].vu == false {
				fmt.Println(tmp)
				fmt.Println("up : ", tmp.kid[i])
				return &tmp.kid[i]
			}
		}
	}

	for i := 0; i < len(solver.kid) ; i++ {
		if solver.kid[i].vu == false {
			fmt.Println("open : ", solver)
			fmt.Println("out" ,solver.kid[i], solver.kid[i].vu)
			return &solver.kid[i]
		}
	}

	if solver.dady != nil {
		fmt.Println("out dady" ,*solver.dady, (*solver.dady).vu)
		return solver.dady
	}
	fmt.Println("fail: ", solver)
//	solver.state.score = -1
	return solver
}

func uniformCostSearch(solver Solver) (e error) {
	explored = make(map[string]bool, 10000)
	count = 0
	depth := 0
	nb_open = 0
	fmt.Println(solver)
	for solver.state.score != 0 && solver.state.score != -1{
		solver = *appendAllNewStates(&solver, depth)
		fmt.Println("call:", solver, solver.vu)
		appendAllNewStates(&solver, depth)
	}
	fmt.Println("exit?", solver.depth, solver.dady)
	for solver.dady != nil {
		fmt.Println(solver)
		solver = *solver.dady
	}
	/*	//for solver.Solution[len(solver.Solution)-1].score > 0 && len(solver.openedStates) > 0 {
			solver.appendAllNewStates(0)
			if len(solver.Solution) == 0 {
				return errors.New("ERROR: Unsolvable puzzle 1")
			}
		//}
		if len(solver.Solution) == 0 {
			return errors.New("ERROR: Unsolvable puzzle 2")
		}  */
	return
}
