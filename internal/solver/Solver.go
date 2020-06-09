package solver

import (
	"fmt"
	"time"

	"github.com/Karocyt/Npupu/internal/sortedhashedtree"
)

// scoreFn type: heuristic functions prototype
type scoreFn func([]int, int, int) float32

var size int
var goalKey string

type counters struct {
	counter           uint64
	maxStates         uint64
	totalOpenedStates uint64
	totalStates       uint64
	startTime         time.Time
	totalTime         time.Duration
}

// Solver contains all variables required to solve the grid
// Solver.Solution contains ordered states from the starting grid to the solved one
type Solver struct {
	counters
	openedStates *sortedhashedtree.SortedHashedTree
	fn           scoreFn
	depth        int
	Solution     chan []*gridState
	E            error
	Stats        chan counters
}

// Init initialize globals
func Init(gridSize int) {
	size = gridSize
	goalKey = makeGoalKey(size)
}

// New initialize a new solverStruct, required to disciminate variables in multi-solving
func New(grid []int, gridSize int, fn scoreFn) *Solver {
	solver := Solver{
		counters:     counters{},
		fn:           fn,
		openedStates: sortedhashedtree.New(),
		Solution:     make(chan []*gridState, 1),
		Stats:        make(chan counters, 1),
	}

	state := newGrid(nil, &solver.counter)
	state.path = make([]*gridState, 0)
	state.grid = grid
	state.depth = 1
	state.score = fn(grid, gridSize, 1)
	state.key = state.mapKey()

	solver.AppendState(&state)
	return &solver
}

// PrintStats does exactly what it says
func PrintStats(stats counters) {
	fmt.Println("Total time elapsed:", stats.totalTime)
	fmt.Printf("Total states analyzed: %d\n", stats.totalStates)
	fmt.Printf("Total states selected: %d\n", stats.totalOpenedStates)
	fmt.Printf("Maximum states ever represented at once: %d\n\n", stats.maxStates)
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
func (solver *Solver) PrintRes(name string, solution []*gridState, success bool, stats counters, display bool) {
	if display {
		for _, step := range solution {
			fmt.Println(step)
		}
	}
	fmt.Printf("Solution using %s:\n\n", name)
	if success {
		fmt.Printf("Solution found: %d moves\n", len(solution))
	} else {
		fmt.Println("This puzzle is not solvable")
	}
	PrintStats(stats)
}

// AppendState prout
func (solver *Solver) AppendState(state *gridState) bool {
	return solver.openedStates.Insert(state.key, state, state.score)
}
