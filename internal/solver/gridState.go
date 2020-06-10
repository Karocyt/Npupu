package solver

import (
	"fmt"
)

// gridState type: grid format/interface
type gridState struct {
	grid  []int
	depth int
	score float32
	path  []int8
	key   string
	id    int8
}

func get1d(x int, y int, size int) int {
	return x*size + y
}

func get2d(nb int, size int) (int, int) {
	var x, y int
	x = nb / size
	y = nb % size
	return x, y
}

func (state *gridState) getVal(x, y int) int {
	return state.grid[x*size+y]
}

func (state *gridState) setVal(x, y, value int) {
	state.grid[x*size+y] = value
}

func (state *gridState) generateState(xZero, yZero, xTarget, yTarget int) *gridState {
	newState := newGrid(state)
	newState.path = make([]int8, len(state.path)+1)
	copy(newState.path, state.path)
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
	for i := range ret {
		ret[i].path[len(state.path)] = int8(i)
	}
	return ret
}

// gridState Stringify function
func (state gridState) String() string {
	nbSize := len(fmt.Sprintf("%d", size*size-1))
	nbWall := "###"
	for i := 0; i < nbSize; i++ {
		nbWall += "#"
	}
	wall := ""
	for i := 0; i < size; i++ {
		wall += nbWall
	}
	wall += "#\n"
	s := wall
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if state.grid[i*size+j] != 0 {
				s += fmt.Sprintf("# %*d ", nbSize, state.grid[i*size+j])
			} else {
				s += fmt.Sprintf("# %*s ", nbSize, "")
			}
		}
		s += "#\n"
		s += wall
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
	var n gridState
	return n
}
