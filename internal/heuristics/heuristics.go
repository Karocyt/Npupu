package heuristics

// all heuristic functions should be of type func([]int) string

type heuristicFn struct {
	Fn   func(grid []int) int
	Name string
}

// Functions is our slice of heuristics as couples {function pointer, name}
var Functions = []heuristicFn{
	heuristicFn{test, "Fake heuristic"},
}

// test function: most basic heuristic we can do, scoring nothing and returning 0, yay!
func test(grid []int) int {
	return (0)
}
