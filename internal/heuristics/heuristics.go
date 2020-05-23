package heuristics

// all heuristic functions should be of type func([]int) string

var nbPos map[int] [2]int


type heuristicFn struct {
	Fn     func(grid []int, size int, depth int) float32
	Name   string
	Greedy bool
}

// Functions is our slice of heuristics as couples {function func, name string, greedy bool}
var Functions = []heuristicFn{
	heuristicFn{toop, "Fake heuristic", true},
	heuristicFn{manhattan, "Greedy Manhattan Distance", true},
	heuristicFn{manhattan, "Uniform-cost Manhattan Distance", false},
	heuristicFn{euclidean, "Greedy Euclidean Distance", true},
	heuristicFn{euclidean, "Uniform-cost Euclidean Distance", false},
}

// test function: Dummy heuristic, scoring nothing and returning 0, yay!
func test(grid []int, size int) float32 {
	return (0)
}





