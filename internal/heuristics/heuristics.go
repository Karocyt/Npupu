package heuristics

// all heuristic functions should be of type func([]int) string

var nbPos map[int][2]int

type heuristicFn struct {
	Fn   func(grid []int, size int, depth int) float32
	Name string
}

// Functions is our slice of heuristics as couples {function func, name string, greedy bool}
var Functions = []heuristicFn{
	heuristicFn{toop, "Fake heuristic"},
	heuristicFn{uniform, "Uniform-cost"},
	heuristicFn{manhattan, "Greedy Manhattan Distance"},
	heuristicFn{manhattanA, "A* Manhattan Distance"},
	heuristicFn{euclidean, "Greedy Euclidean Distance"},
	heuristicFn{euclideanA, "A* Euclidean Distance"},
}

// test function: Dummy heuristic, scoring nothing and returning 0, yay!
func test(grid []int, size int, depth int) float32 {
	return (0)
}
