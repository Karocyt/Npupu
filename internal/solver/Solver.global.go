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

// Solve solve
func (solver *Solver) Solve() {
	solver.startTime = time.Now()
	cur := solver.openedStates.GetMin().(*gridState)
	root := cur
	for cur != nil && cur.key != goalKey {
		nextStates := cur.generateNextStates()
		var included int
		for i := range nextStates {
			if solver.openedStates.IsInHistory(nextStates[i].key) == false {
				nextStates[i].score = solver.fn(nextStates[i].grid, size, nextStates[i].depth)
				solver.AppendState(nextStates[i])
				included++
			}
		}
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
		ret := make([]*gridState, 1, 50)
		ret[0] = root
		for i, id := range cur.path {
			step := ret[i].generateNextStates()[id]
			step.score = solver.fn(step.grid, size, step.depth)
			ret = append(ret, step)
		}
		solver.Solution <- ret
	}
	solver.totalTime = time.Since(solver.startTime)
	solver.counters.totalStates, solver.counters.totalOpenedStates, solver.counters.maxStates = solver.openedStates.GetStats()
	solver.Stats <- solver.counters
	return
}
