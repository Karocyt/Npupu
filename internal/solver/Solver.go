package solver

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Karocyt/Npupu/internal/sortedhashedtree"
)

// ScoreFn type: heuristic functions prototype
type ScoreFn func([]int, int, int) float32

var size int
var goalKey string
var goalMap map[int][2]int
var finalGrid []int

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
func Init(gridSize int, classic bool, input []int, randomSize int, shuffleCount int) (map[int][2]int, []int, []int) {
	rand.Seed(time.Now().UnixNano())
	size = gridSize
	if input == nil {
		size = randomSize
	}
	goalKey, goalMap, finalGrid = makeGoalState(classic)
	if input == nil {
		input = pupuRand(shuffleCount)
	} else if !checkSolvy(input, classic) {
		fmt.Println("Pupu is not solvable :3")
		os.Exit(0)
	}
	return goalMap, finalGrid, input
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
	fmt.Printf("Total states generated (time complexity): %d\n", stats.totalStates)
	fmt.Printf("Total states expanded: %d\n", stats.totalOpenedStates)
	fmt.Printf("Maximum states in the open set at once (space complexity): %d\n\n", stats.maxStates)
}

func makeGoalState(classic bool) (string, map[int][2]int, []int) {
	s := size
	nbPos := make(map[int][2]int)
	puzzle := make([]int, s*s)
	cur := 1
	x := 0
	ix := 1
	y := 0
	iy := 0
	if classic {
		for cur < s*s {
			if x == s {
				y++
				x = 0
			}
			puzzle[x+y*s] = cur
			nbPos[cur] = [2]int{y, x}
			cur++

			x++
		}
		nbPos[0] = [2]int{y, x}
		puzzle[x+y*s] = 0
	} else {
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
	}

	return gridState{grid: puzzle}.mapKey(), nbPos, puzzle
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
		fmt.Printf("Solution found: %d moves\n", len(solution)-1)
	} else {
		fmt.Println("This puzzle is not solvable")
	}
	PrintStats(stats)
}

// AppendState prout
func (solver *Solver) AppendState(state *gridState) bool {
	return solver.openedStates.Insert(state.key, state, state.score)
}
