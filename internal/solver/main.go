package solver

// scoreFn type: heuristic functions prototype
type scoreFn func([]int) int

// New initialize a new solverStruct, required to disciminate variables in multi-solving
// Can be removed if we don't need to initialize anything
// (we can use "var s Solver.Solver" in main instead of calling this)
func New() Solver {
	var solver Solver
	return solver
}
