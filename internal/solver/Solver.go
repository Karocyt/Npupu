package solver

import "errors"

// Solver contains all variables required to solve the grid
// Solver.Solution contains ordered states from the starting grid to the solved one
type Solver struct {
	openedStatesCount int
	maxOpenedStates   int
	openedStates      []gridState
	depth             int
	explored          map[string]bool
	Solution          []gridState
	fn                scoreFn
}

func (solver *Solver) hasSeen(state gridState) bool {
	key := state.String() // might be better with a fastString() method
	return solver.explored[key]
}

func (solver *Solver) appendNextStates() {
	state := solver.openedStates[len(solver.openedStates)-1]
	/* TODO:
	if state.depth != solver.depth
		pop from solution and decrement solver.depth
	else
		increment solver.depth
	// might need to do something with openStatesCount
	*/
	key := state.String()
	solver.explored[key] = true
	for _, newState := range state.generateNextStates() {
		if solver.hasSeen(newState) == false {
			newState.score = solver.fn(newState.grid)
			solver.openedStates = append(solver.openedStates, newState)
			solver.openedStatesCount++
		}
	}
}

/*Solve Function:
** Solves a given puzzle with A* algorithm.
** 	- 1st argument: a grid in the format [N*N]int
** 	- 2nd argument: size N of the aforementioned grid
** 	- 3rd argument: score function of type 'func([]int) int' used as heuristic
** return value: error e (unsolvable)
 */
func (solver *Solver) Solve(grid []int, gridSize int, fn scoreFn) (e error) {
	state := gridState{
		grid:  grid,
		size:  gridSize,
		depth: 0,
		score: fn(grid),
	}
	solver.fn = fn
	solver.Solution = append(solver.Solution, state)
	solver.explored = make(map[string]bool, 1)
	solver.openedStatesCount++

	for solver.Solution[len(solver.Solution)-1].score > 0 {
		solver.appendNextStates()
		if len(solver.Solution) == 0 {
			return errors.New("ERROR: Unsolvable puzzle")
		}
	}
	return
}
