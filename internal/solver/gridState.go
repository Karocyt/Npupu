package solver

import "fmt"

// gridState type: grid format/interface
type gridState struct {
	grid  []int
	size  int // to avoid using math.Sqrt everywhere
	depth int
	score int
}

func (state *gridState) get(x, y int) int {
	return state.grid[x*state.size+y]
}

func (state *gridState) set(x, y, value int) {
	state.grid[x*state.size+y] = value
}

func (state *gridState) generateNextStates() []gridState {
	/* TO DO */

	return []gridState{}
}

// gridState Stringify function
func (state gridState) String() string {
	var s string
	for i := 0; i < state.size; i++ {
		s += "#####"
	}
	s += "\n"
	for i := 0; i < state.size; i++ {
		for j := 0; j < state.size; j++ {
			s += fmt.Sprintf("#%2d #", state.grid[i*state.size+j])
		}
		s += "\n"
		for j := 0; j < state.size; j++ {
			s += "#####"
		}
		s += "\n"
	}
	return fmt.Sprintf("Step %d:\n%sScore: %d\n", state.depth, s, state.score)
}
