package solver

import (
	"time"
)

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
	solver.startTime = time.Now()
	cur := solver.openedStates.GetMin().(*gridState)
	for cur != nil && cur.key != goalKey {
		nextStates := cur.generateNextStates(&solver.counter)

		var included int
		for i := range nextStates {
			if solver.openedStates.IsInHistory(nextStates[i].key) == false {
				nextStates[i].score = solver.fn(nextStates[i].grid, size, nextStates[i].depth)
				solver.AppendState(nextStates[i])
				included++
			}
		}
		solver.counter -= uint64(len(nextStates) - included + 1)
		solver.decrementParents(cur)
		solver.openedStates.Delete(cur.key)
		if cur.key != goalKey {
			tmp := solver.openedStates.GetMin()
			if tmp != nil {
				cur = tmp.(*gridState)
			} else {
				cur = nil
			}
		}
	}
	if solver.openedStates.GetLen() == 0 {
		close(solver.Solution)
	} else {
		solver.Solution <- append(cur.path, cur)
	}
	solver.totalTime = time.Since(solver.startTime)
	solver.counters.totalStates, solver.counters.totalOpenedStates, solver.counters.maxStates = solver.openedStates.GetStats()
	solver.Stats <- solver.counters
	return
}
