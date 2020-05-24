package solver

func bestScore(l []*gridState) (cur *gridState) {
	for _, item := range l {
		if cur == nil || item.score < cur.score {
			cur = item
		}
	}
	return cur
}

func (solver *Solver) decrementParents(state *gridState) {
	for i := range state.path {
		state.path[i].childsCount--
		if state.path[i].childsCount == 0 {
			solver.counter--
		}
	}
}

// Solve solve
func (solver *Solver) Solve() {
	cur := solver.openedStates.PopMin().Value.(*gridState)
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
				solver.AppendState(nextStates[i])
				solver.totalStates++
				included++
			}
		}
		if solver.counter > solver.maxStates {
			solver.maxStates = solver.counter
		}
		solver.counter -= (len(nextStates) - included)
		solver.decrementParents(cur)
		if curKey != goalKey {
			tmp := solver.openedStates.PopMin()
			if tmp != nil {
				cur = tmp.Value.(*gridState)
			} else {
				cur = nil
			}
		}
	}
	if solver.openedStates.GetCount() == 0 {
		close(solver.Solution)
	} else {
		solver.Solution <- append(cur.path, cur)
	}
	solver.Stats <- solver.counters
	return
}
