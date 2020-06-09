package solver

import (
	"fmt"
	"time"

	"github.com/Karocyt/Npupu/internal/sortedhashedtree"
)

// ScoreFn type: heuristic functions prototype
type ScoreFn func([]int, int, int) float32

var size int
var goalKey string

type counters struct {
	maxStates         int
	totalOpenedStates int
	totalStates       int
	startTime         time.Time
	totalTime         time.Duration
}

// Solver contains all variables required to solve the grid
// Solver.Solution contains ordered states from the starting grid to the solved one
type Solver struct {
	counters
	Name         string
	openedStates *sortedhashedtree.SortedHashedTree
	fn           ScoreFn
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
func New(grid []int, gridSize int, fn ScoreFn, name string) *Solver {
	solver := Solver{
		counters:     counters{},
		Name:         name,
		fn:           fn,
		openedStates: sortedhashedtree.New(),
		Solution:     make(chan []*gridState, 1),
		Stats:        make(chan counters, 1),
	}

	state := newGrid(nil)
	state.grid = grid
	state.depth = 0
	state.score = fn(grid, gridSize, 1)
	state.key = state.mapKey()
	state.id = 42
	state.path = []int8{}

	solver.AppendState(&state)
	return &solver
}

// PrintStats does exactly what it says
func PrintStats(stats counters) {
	fmt.Println("Total time elapsed:", stats.totalTime)
	fmt.Printf("Total states generated: %d\n", stats.totalStates)
	fmt.Printf("Total states selected: %d\n", stats.totalOpenedStates)
	fmt.Printf("Maximum states in the open set at once: %d\n", stats.maxStates)
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
func (solver *Solver) PrintRes(solution []*gridState, success bool, stats counters, display bool) {
	if display {
		for _, step := range solution {
			fmt.Println(step)
		}
	}
	fmt.Printf("Solution using %s:\n\n", solver.Name)
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
