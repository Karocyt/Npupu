package solverv2

var count int

func appendAllNewStates(solver *Solver, stop int) {
	if stop == 1 {
		return
	}
	state := solver.openedStates[len(solver.openedStates)-1]
	key := state.mapKey()
	explored[key] = true
	nextStates := make([]gridState, 0, 4)

	for len(nextStates) == 0 && len(solver.openedStates) > 0 {
		state := solver.openedStates[len(solver.openedStates)-1]
		solver.depth = state.depth
		for _, newState := range state.generateNextStates() {
			if hasSeen(newState) == false {
				newState.score = solver.fn(newState.grid, size)
				nextStates = append(nextStates, newState)
			}
		}

	/*	if len(nextStates) == 0 && len(solver.openedStates) > 0 {
			solver.openedStates = solver.openedStates[0 : len(solver.openedStates)-1]
			if len(solver.Solution) > 0 {
				solver.Solution = solver.Solution[0 : len(solver.Solution)-1]
			}
		}
	} */

	/*for i := 0; i < len(nextStates); i++ {
		if nextStates[i].score == 0 {
			stop = 1
		}
		 /*if solver.explored[nextStates[i].mapKey()] == false {

		//	solver.Solution = append(solver.Solution, nextStates[i])
		///	solver.openedStates = append(solver.openedStates, nextStates[i])
			solver.totalOpenedStates++
			solver.depth = nextStates[i].depth
			solver.appendAllNewStates()
		} */
	}

	/*
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

	*/
/*
	if state.score != 0 && solver.explored[nextStates[0].mapKey()] == false && stop == 0{
		if state.score == 0 {
			stop = 1
		}
		solver.appendAllNewStates()
	} */
}

func uniformCostSearch(solver Solver) (e error) {
	explored = make(map[string]bool, 10000)
	count = 0
	//for solver.Solution[len(solver.Solution)-1].score > 0 && len(solver.openedStates) > 0 {
	/*	solver.appendAllNewStates(0)
		if len(solver.Solution) == 0 {
			return errors.New("ERROR: Unsolvable puzzle 1")
		}
	//}
	if len(solver.Solution) == 0 {
		return errors.New("ERROR: Unsolvable puzzle 2")
	} */
	return
}
