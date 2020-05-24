package heuristics

// all heuristic functions should be of type func([]int) string

type heuristicFn struct {
	Fn   func(grid []int, size int, depth int) float32
	Name string
}

// Functions is our slice of heuristics as couples {function func, name string, greedy bool}
var Functions = []heuristicFn{
	heuristicFn{test, "Fake heuristic"},
	heuristicFn{uniform, "Uniform-cost"},
	heuristicFn{manhattan, "Greedy Manhattan Distance"},
	heuristicFn{manhattanA, "A* Manhattan Distance"},
	heuristicFn{euclidean, "Greedy Euclidean Distance"},
	heuristicFn{euclideanA, "A* Euclidean Distance"},
	heuristicFn{toop, "Greedy Tiles-Out-Of-Place"},
	heuristicFn{toopA, "A* Tiles-Out-Of-Place"},
}

// test function: Dummy heuristic, scoring nothing and returning 0, yay!
func test(grid []int, size int, depth int) float32 {
	return (0)
}
