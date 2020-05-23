package solverv2

// scoreFn type: heuristic functions prototype
type scoreFn func([]int, int) float32

var explored          map[string]bool
var totalOpenedStates int

// Solver contains all variables required to solve the grid
// Solver.Solution contains ordered states from the starting grid to the solved one


type Solver struct {
	dady			*Solver
	id 				int
	state			gridState
	depth			int
	fn              scoreFn
	kid 			[]Solver
	open 			bool
	vu 				bool
}

var size int

// New initialize a new solverStruct, required to disciminate variables in multi-solving
// Can be removed if we don't need to initialize anything
// (we can use "var s Solver.Solver" in main instead of calling this)
func New(grid []int, gridSize int, fn scoreFn, greedy bool) Solver {
	var solver Solver

	size = gridSize
	state := gridState{
		grid:  grid,
		depth: 0,
		score: fn(grid, gridSize),
	}
	solver.fn = fn
//	solver.Solution = append(solver.Solution, state)
//	solver.explored = make(map[string]bool, 1000)
	solver.state = state
	totalOpenedStates++
	return solver
}

func hasSeen(state gridState) bool {
	key := state.mapKey()
	return explored[key]
}

/*Solve Function:
** Solves a given puzzle with A* algorithm.
** 	- 1st argument: a grid in the format [N*N]int
** 	- 2nd argument: size N of the aforementioned grid
** 	- 3rd argument: score function of type 'func([]int) int' used as heuristic
** return value: error e (unsolvable)
 */

func Solve(solver Solver) (e error) {
		e = uniformCostSearch(solver)
	return
}
