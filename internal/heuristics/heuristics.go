package heuristics

// all heuristic functions should be of type func([]int) string

type heuristicFn struct {
	Fn    func(grid []int, size int, depth int) float32
	Name  string
	Astar func(grid []int, size int, depth int) float32
}

// Functions is our slice of heuristics as couples {function func, name string, greedy bool}
var Functions = []heuristicFn{
	heuristicFn{uniform, "Uniform-cost", nil},
	heuristicFn{manhattan, "Manhattan Distance", manhattanA},
	heuristicFn{euclidean, "Euclidean Distance", euclideanA},
	heuristicFn{toop, "Tiles-Out-Of-Place", toopA},
}
