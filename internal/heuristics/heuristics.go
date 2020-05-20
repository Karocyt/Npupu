package heuristics

// all heuristic functions should be of type func([]int) string

type heuristicFn struct {
	Fn     func(grid []int, size int) int
	Name   string
	Greedy bool
}

// Functions is our slice of heuristics as couples {function func, name string, greedy bool}
var Functions = []heuristicFn{
	heuristicFn{test, "Fake heuristic", true},
	heuristicFn{manhattan, "Greedy Manhattan Distance", true},
	heuristicFn{manhattan, "Uniform-cost Manhattan Distance", false},
	heuristicFn{euclidean, "Greedy Euclidean Distance", true},
	heuristicFn{euclidean, "Uniform-cost Euclidean Distance", false},
}

// test function: Dummy heuristic, scoring nothing and returning 0, yay!
func test(grid []int, size int) int {
	return (0)
}

// manhattan distance function: most basic heuristic we can do
func manhattan(grid []int, size int) int {
	return (0)
}

// euclidean distance function: other basic heuristic
func euclidean(grid []int, size int) int {
	return (0)
}
