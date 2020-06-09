package solver

import (
	"fmt"
)

// gridState type: grid format/interface
type gridState struct {
	grid        []int
	depth       int
	score       float32
	parent      *gridState
	childsCount int
	key         string
}

func (state *gridState) getVal(x, y int) int {
	return state.grid[x*size+y]
}

func (state *gridState) setVal(x, y, value int) {
	state.grid[x*size+y] = value
}

func (state *gridState) generateState(xZero, yZero, xTarget, yTarget int) *gridState {
	newState := newGrid(state)
	newState.parent = state
	newState.grid = make([]int, len(state.grid))
	copy(newState.grid, state.grid)
	newState.depth = state.depth + 1
	newState.setVal(xZero, yZero, state.getVal(xTarget, yTarget))
	newState.setVal(xTarget, yTarget, 0)
	newState.key = newState.mapKey()
	return &newState
}

func (state *gridState) generateNextStates() []*gridState {
	ret := make([]*gridState, 0, 4)
	idx := -1
	for i, nb := range state.grid {
		if nb == 0 {
			idx = i
			break
		}
	}
	x, y := idx/size, idx%size
	if x > 0 {
		ret = append(ret, state.generateState(x, y, x-1, y))
	}
	if x < size-1 {
		ret = append(ret, state.generateState(x, y, x+1, y))
	}
	if y > 0 {
		ret = append(ret, state.generateState(x, y, x, y-1))
	}
	if y < size-1 {
		ret = append(ret, state.generateState(x, y, x, y+1))
	}
	return ret
}

// gridState Stringify function
func (state gridState) String() string {
	var s string
	for i := 0; i < size; i++ {
		s += "#####"
	}
	s += "\n"
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if state.grid[i*size+j] != 0 {
				s += fmt.Sprintf("#%2d #", state.grid[i*size+j])
			} else {
				s += "#   #"
			}
		}
		s += "\n"
		for j := 0; j < size; j++ {
			s += "#####"
		}
		s += "\n"
	}
	return fmt.Sprintf("Step %d:\n%sScore: %f\n", state.depth, s, state.score)
}

// mapKey Stringify function to provide keys for our visited map
func (state gridState) mapKey() string {
	s := make([]rune, size*size)
	for i := 0; i < size*size; i++ {
		s[i] = rune(state.grid[i] + 1)
	}
	return string(s)
}

// NewGrid creates a new gridState and manage the states counter
func newGrid(parent *gridState) gridState {
	if parent != nil {
		root := parent
		for root.parent != nil {
			root.childsCount++
			root = root.parent
		}
		root.childsCount++
	}
	var n gridState
	n.parent = parent
	return n
}
