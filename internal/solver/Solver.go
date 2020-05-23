package solver

// scoreFn type: heuristic functions prototype
type scoreFn func([]int, int, int) float32

var size int

// Solver contains all variables required to solve the grid
// Solver.Solution contains ordered states from the starting grid to the solved one
type Solver struct {
	openedStates    []*gridState
	Solution        []*gridState
	fn              scoreFn
	explored        map[string]bool
	maxOpenedStates int
	totalStates     int
	depth           int
}

// New initialize a new solverStruct, required to disciminate variables in multi-solving
// Can be removed if we don't need to initialize anything
// (we can use "var s Solver.Solver" in main instead of calling this)
func New(grid []int, gridSize int, fn scoreFn, greedy bool) Solver {
	solver := Solver{
		totalStates: 1,
		fn:          fn,
		explored:    make(map[string]bool, 100*size*size),
	}

	size = gridSize
	state := gridState{
		grid:  grid,
		depth: 1,
		score: fn(grid, gridSize, 1),
	}

	solver.openedStates = append(solver.openedStates, &state)

	return solver
}

func (solver *Solver) hasSeen(state gridState) bool {
	key := state.mapKey()
	return solver.explored[key]
}

// PrintStats does exactly what it says
func (solver *Solver) PrintStats() {
	/* TODO */
}
