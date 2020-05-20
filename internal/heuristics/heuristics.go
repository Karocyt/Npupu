package heuristics

// all heuristic functions should be of type func([]int) string

type heuristicFn struct {
	Fn      func(grid []int) int
	Name    string
	Uniform bool
}

// Functions is our slice of heuristics as couples {function func, name string, greedy bool}
var Functions = []heuristicFn{
	heuristicFn{test, "Fake heuristic", true},
	heuristicFn{test, "Greedy Manhattan Distance", true},
	heuristicFn{test, "Uniform-cost Manhattan Distance", false},
}

// test function: most basic heuristic we can do, scoring nothing and returning 0, yay!
func test(grid []int) int {
	return (0)
}
