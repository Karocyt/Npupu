package solver

import (
	"fmt"
)

// gridState type: grid format/interface
type gridState struct {
	grid        []int
	depth       int
	score       float32
	path        []*gridState
	childsCount int
}

func (state *gridState) getVal(x, y int) int {
	return state.grid[x*size+y]
}

func (state *gridState) setVal(x, y, value int) {
	state.grid[x*size+y] = value
}

func (state *gridState) generateState(xZero, yZero, xTarget, yTarget int, counter *int) *gridState {
	newPath := make([]*gridState, len(state.path), len(state.path)+1)
	copy(newPath, state.path)
	newPath = append(newPath, state)
	newState := newGrid(newPath, counter)
	newState.path = append(state.path, state)
	newState.grid = make([]int, len(state.grid))
	copy(newState.grid, state.grid)
	newState.depth = state.depth + 1
	newState.setVal(xZero, yZero, state.getVal(xTarget, yTarget))
	newState.setVal(xTarget, yTarget, 0)
	return &newState
}

func (state *gridState) generateNextStates(counter *int) []*gridState {
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
		ret = append(ret, state.generateState(x, y, x-1, y, counter))
	}
	if x < size-1 {
		ret = append(ret, state.generateState(x, y, x+1, y, counter))
	}
	if y > 0 {
		ret = append(ret, state.generateState(x, y, x, y-1, counter))
	}
	if y < size-1 {
		ret = append(ret, state.generateState(x, y, x, y+1, counter))
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
func newGrid(path []*gridState, counter *int) gridState {
	for i := range path {
		path[i].childsCount++
	}
	(*counter)++
	var n gridState
	n.path = make([]*gridState, len(path))
	copy(n.path, path)
	return n
}
