package solver

import (
	"errors"
	"fmt"
	"sort"
)

func (solver *Solver) appendNextStates() {
	state := solver.openedStates[len(solver.openedStates)-1]
	key := state.mapKey()
	solver.explored[key] = true
	nextStates := make([]gridState, 0, 4)
	for len(nextStates) == 0 && len(solver.openedStates) > 0 {
		solver.openedStates = solver.openedStates[0 : len(solver.openedStates)-1]
		if len(solver.openedStates) > 1 {
			state := solver.openedStates[len(solver.openedStates)-1]
			solver.depth = state.depth
			for i, newState := range state.generateNextStates() {
				if solver.hasSeen(newState) == false {
					nextStates = append(nextStates, newState)
					nextStates[i].score = solver.fn(nextStates[i].grid, size)
				}
			}
		}
	}
	sort.Slice(nextStates, func(i, j int) bool {
		return nextStates[i].score > nextStates[j].score
	})
	fmt.Println(state)
	for _, newState := range nextStates {
		if solver.hasSeen(newState) == false {
			fmt.Println(newState.score)
			solver.openedStates = append(solver.openedStates, newState)
			solver.totalOpenedStates++
			solver.depth = newState.depth
		}
	}
	fmt.Println()
	if len(solver.openedStates) > solver.maxOpenedStates {
		solver.maxOpenedStates = len(solver.openedStates)
	}
}

func (solver *Solver) greedySearch() (e error) {
	for solver.Solution[len(solver.Solution)-1].score > 0 && len(solver.openedStates) > 0 {
		solver.appendNextStates()
		if len(solver.Solution) == 0 {
			return errors.New("ERROR: Unsolvable puzzle 1")
		}
	}
	if len(solver.Solution) == 0 {
		return errors.New("ERROR: Unsolvable puzzle 2")
	}
	return
}
