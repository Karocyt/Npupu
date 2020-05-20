package solver

// scoreFn type: heuristic functions prototype
type scoreFn func([]int, int) int

// Solver contains all variables required to solve the grid
// Solver.Solution contains ordered states from the starting grid to the solved one
type Solver struct {
	maxOpenedStates   int
	totalOpenedStates int
	openedStates      []gridState
	depth             int
	explored          map[string]bool
	Solution          []gridState
	fn                scoreFn
	greedy            bool
}

// New initialize a new solverStruct, required to disciminate variables in multi-solving
// Can be removed if we don't need to initialize anything
// (we can use "var s Solver.Solver" in main instead of calling this)
func New(grid []int, gridSize int, fn scoreFn, greedy bool) Solver {
	var solver Solver

	state := gridState{
		grid:  grid,
		size:  gridSize,
		depth: 0,
		score: fn(grid, gridSize),
	}
	solver.fn = fn
	solver.Solution = append(solver.Solution, state)
	solver.explored = make(map[string]bool, 1000)
	solver.openedStates = append(solver.openedStates, state)
	solver.totalOpenedStates++
	solver.greedy = greedy
	return solver
}

func (solver *Solver) hasSeen(state gridState) bool {
	key := state.String() // might be better with a fastString() method
	return solver.explored[key]
}

/*Solve Function:
** Solves a given puzzle with A* algorithm.
** 	- 1st argument: a grid in the format [N*N]int
** 	- 2nd argument: size N of the aforementioned grid
** 	- 3rd argument: score function of type 'func([]int) int' used as heuristic
** return value: error e (unsolvable)
 */
func (solver *Solver) Solve() (e error) {
	if solver.greedy {
		e = solver.greedySearch()
	} else {
		e = solver.uniformCostSearch()
	}
	return
}

// PrintStats does exactly what it says
func (solver *Solver) PrintStats() {
	/* TODO */
}
