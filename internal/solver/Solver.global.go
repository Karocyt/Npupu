package solver

import (
	"fmt"
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

func (solver *Solver) decrementCount(state *gridState) {
	if state.childsCount == 0 {
		solver.counter--
	}
	for state.parent != nil {
		state = state.parent
		state.childsCount--
		if state.childsCount == 0 {
			solver.counter--
		}
	}
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
				solver.counter++
				nextStates[i].score = solver.fn(nextStates[i].grid, size, nextStates[i].depth)
				solver.AppendState(nextStates[i])
				included++
			}
		}
		// if uint64(root.childsCount) >= solver.worstCase {
		// 	solver.worstCase = uint64(root.childsCount) + 1
		// }
		solver.decrementCount(cur)
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
		ret := make([]*gridState, 0, 50)
		for cur != nil {
			ret = append(ret, cur)
			cur = cur.parent
		}
		for i := len(ret)/2 - 1; i >= 0; i-- {
			opp := len(ret) - 1 - i
			ret[i], ret[opp] = ret[opp], ret[i]
		}
		solver.Solution <- ret
	}
	fmt.Println(root.childsCount)
	solver.totalTime = time.Since(solver.startTime)
	solver.counters.totalStates, solver.counters.totalOpenedStates, solver.counters.maxStates = solver.openedStates.GetStats()
	solver.Stats <- solver.counters
	return
}
