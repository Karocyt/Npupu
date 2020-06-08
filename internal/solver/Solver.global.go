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
	curKey := cur.mapKey()
	// fmt.Println("0", cur.score)
	for cur != nil && curKey != goalKey {
		curKey = cur.mapKey()
		// fmt.Printf("%p\n", cur)
		nextStates := cur.generateNextStates(&solver.counter)
		solver.explored[curKey] = true
		solver.totalOpenedStates++

		// fmt.Println("1 (", len(nextStates), ")")
		var included int
		for i := range nextStates {
			if solver.explored[nextStates[i].mapKey()] == false {
				nextStates[i].score = solver.fn(nextStates[i].grid, size, nextStates[i].depth)
				// fmt.Println("1.", i, nextStates[i].score)
				//fmt.Println("/thead score:", i, solver.openedStates.GetMin().(*gridState).score)
				solver.AppendState(nextStates[i])
				solver.totalStates++
				included++
			}
		}
		// fmt.Println("2")
		if solver.counter > solver.maxStates {
			solver.maxStates = solver.counter
		}
		solver.counter -= (len(nextStates) - included + 1)
		solver.decrementParents(cur)
		// fmt.Println(curKey)
		solver.openedStates.Delete(curKey)
		// fmt.Println("3", cur.score)
		if curKey != goalKey {
			tmp := solver.openedStates.GetMin()
			if tmp != nil {
				cur = tmp.(*gridState)
			} else {
				cur = nil
			}
		}
	}
	//fmt.Println("YOUPI")
	if solver.openedStates.GetLen() == 0 {
		close(solver.Solution)
	} else {
		solver.Solution <- append(cur.path, cur)
	}
	solver.totalTime = time.Since(solver.startTime)
	solver.Stats <- solver.counters
	return
}
