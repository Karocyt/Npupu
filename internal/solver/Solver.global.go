package solver

func bestScore(l []*gridState) (cur *gridState) {
	for _, item := range l {
		if cur == nil || item.score < cur.score {
			cur = item
		}
	}
	return cur
}

func (solver *Solver) findState(state *gridState) int {
	for i, find := range solver.openedStates {
		if find == state {
			return i
		}
	}
	return -1
}

func (solver *Solver) closeState(state *gridState) {
	idx := solver.findState(state)
	solver.openedStates[idx] = solver.openedStates[len(solver.openedStates)-1]
	solver.openedStates[len(solver.openedStates)-1] = nil // To remove ?
	solver.openedStates = solver.openedStates[:len(solver.openedStates)-1]
	for i := range state.path {
		state.path[i].childsCount--
		if state.path[i].childsCount == 0 {
			solver.counter--
		}
	}
}

// Solve solve
func (solver *Solver) Solve() {
	cur := solver.openedStates[0]
	curKey := cur.mapKey()
	for cur != nil && curKey != goalKey {
		curKey = cur.mapKey()
		nextStates := cur.generateNextStates(&solver.counter)
		solver.explored[curKey] = true
		solver.totalOpenedStates++

		var included int
		for i := range nextStates {
			if solver.explored[nextStates[i].mapKey()] == false {
				nextStates[i].score = solver.fn(nextStates[i].grid, size, nextStates[i].depth)
				solver.openedStates = append(solver.openedStates, nextStates[i])
				solver.totalStates++
				included++
			}
		}
		if solver.counter > solver.maxStates {
			solver.maxStates = solver.counter
		}
		solver.counter -= (len(nextStates) - included)
		solver.closeState(cur)
		if curKey != goalKey {
			cur = bestScore(solver.openedStates)
		}
	}
	if len(solver.openedStates) == 0 {
		close(solver.Solution)
	} else {
		solver.Solution <- append(cur.path, cur)
	}
	solver.Stats <- solver.counters
	return
}
