package heuristics

// all heuristic functions should be of type func([]int, int, int) string

type heuristicFn struct {
	Fn    func(grid []int, size int, depth int) float32
	Name  string
	Astar func(grid []int, size int, depth int) float32
}

var finalState []int
var finalPos map[int][2]int

// Functions is our slice of heuristics as couples {score func, name string, aStar func}
var Functions = []heuristicFn{
	heuristicFn{uniform, "Uniform-cost", nil},
	heuristicFn{manhattan, "Manhattan Distance", manhattanA},
	heuristicFn{euclidean, "Euclidean Distance", euclideanA},
	heuristicFn{toop, "Tiles-Out-Of-Place", toopA},
}

// Init sets global vars
func Init(grid []int, pos map[int][2]int) {
	finalState = grid
	finalPos = pos
}
