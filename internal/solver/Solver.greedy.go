package solver

import (
	"errors"
	"sort"
)

func (solver *Solver) appendNextStates() {
	state := solver.openedStates[len(solver.openedStates)-1]
	/* TODO:
	if state.depth != solver.depth
		pop from solution and decrement solver.depth
	else
		increment solver.depth
	*/
	key := state.String()
	solver.explored[key] = true
	nextStates := state.generateNextStates()
	sort.Slice(nextStates, func(i, j int) bool {
		return nextStates[i].score < nextStates[j].score
	})
	for _, newState := range nextStates {
		if solver.hasSeen(newState) == false {
			newState.score = solver.fn(newState.grid, newState.size)
			solver.openedStates = append(solver.openedStates, newState)
			solver.totalOpenedStates++
		}
	}
	if len(solver.openedStates) > solver.maxOpenedStates {
		solver.maxOpenedStates = len(solver.openedStates)
	}
}

func (solver *Solver) greedySearch() (e error) {
	for solver.Solution[len(solver.Solution)-1].score > 0 {
		solver.appendNextStates()
		if len(solver.Solution) == 0 {
			return errors.New("ERROR: Unsolvable puzzle")
		}
	}
	return
}
