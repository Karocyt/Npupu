package solver

import "fmt"

// scoreFn type: heuristic functions prototype
type scoreFn func([]int, int, int) float32

var size int
var goalKey string

// Solver contains all variables required to solve the grid
// Solver.Solution contains ordered states from the starting grid to the solved one
type Solver struct {
	openedStates      []*gridState
	Solution          []*gridState
	fn                scoreFn
	explored          map[string]bool
	maxStates         int
	totalOpenedStates int
	totalStates       int
	depth             int
}

// New initialize a new solverStruct, required to disciminate variables in multi-solving
// Can be removed if we don't need to initialize anything
// (we can use "var s Solver.Solver" in main instead of calling this)
func New(grid []int, gridSize int, fn scoreFn) Solver {
	solver := Solver{
		fn:                fn,
		explored:          make(map[string]bool, 100*size*size),
		totalOpenedStates: 0,
		totalStates:       1,
		maxStates:         1,
	}

	size = gridSize
	goalKey = makeGoalKey(size)
	state := gridState{
		grid:  grid,
		depth: 1,
		score: fn(grid, gridSize, 1),
	}

	solver.openedStates = append(solver.openedStates, &state)

	return solver
}

func (solver *Solver) hasSeen(state gridState) bool {
	key := state.mapKey()
	return solver.explored[key]
}

// PrintStats does exactly what it says
func (solver *Solver) PrintStats() {
	fmt.Printf("Total states analyzed: %d\n", solver.totalStates)
	fmt.Printf("Total states selected: %d\n", solver.totalOpenedStates)
	fmt.Printf("Maximum states ever represented at once: %d\n", solver.maxStates)
}

func makeGoalKey(s int) string {
	nbPos := make(map[int][2]int)
	puzzle := make([]int, s*s)
	cur := 1
	x := 0
	ix := 1
	y := 0
	iy := 0
	for cur < s*s {
		puzzle[x+y*s] = cur
		nbPos[cur] = [2]int{y, x}
		cur++

		if x+ix == s || x+ix < 0 || (ix != 0 && puzzle[x+ix+y*s] != 0) {
			iy = ix
			ix = 0
		} else if y+iy == s || y+iy < 0 || (iy != 0 && puzzle[x+(y+iy)*s] != 0) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
	}
	nbPos[0] = [2]int{y, x}
	puzzle[x+y*s] = 0
	grid := gridState{
		grid: puzzle,
	}
	return grid.mapKey()
}

// PrintRes prints.
func (solver *Solver) PrintRes(name string) {
	fmt.Printf("Solution using %s:\n\n", name)
	for _, step := range solver.Solution {
		fmt.Println(step)
	}
	solver.PrintStats()
}
