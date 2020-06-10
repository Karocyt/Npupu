package heuristics

func linearConflicts(grid []int, size int, depth int) float32 {
	conflicts := 0
	for x1 := 0; x1 < size; x1++ {
		for y1 := 0; y1 < size; y1++ {
			if grid[Get1d(x1, y1, size)] != 0 {
				tmp := finalPos[grid[Get1d(x1, y1, size)]]
				targetx, targety := tmp[0], tmp[1]
				if (x1 == targetx) && (y1 != targety) {
					// Case 1: my x is ok
					var incr int
					if targety > y1 {
						incr = 1
					} else {
						incr = -1
					}
					for j := y1 + incr; j != targety+incr; j += incr {
						if grid[Get1d(x1, j, size)] != 0 && finalPos[grid[Get1d(x1, y1, size)]][0] == x1 {
							conflicts++
						}
					}
				} else if (y1 == targety) && (x1 != targetx) {
					// Case 2 my y is ok
					var incr int
					if targetx > x1 {
						incr = 1
					} else {
						incr = -1
					}
					for i := x1 + incr; i != targetx+incr; i += incr {
						if grid[Get1d(i, y1, size)] != 0 && finalPos[grid[Get1d(x1, y1, size)]][1] == y1 {
							conflicts++
						}
					}
				}
			}

		}
	}
	return manhattan(grid, size, depth) + 2*float32(conflicts)
}

// Solution using A* Linear Conflicts:

// Solution found: 58 moves
// Total time elapsed: 12.17546976s
// Total states generated: 5744007
// Total states selected: 2906101
// Maximum states in the open set at once: 2837908

// Solution using A* Manhattan Distance:

// Solution found: 54 moves
// Total time elapsed: 21.983975547s
// Total states generated: 11408118
// Total states selected: 6128715
// Maximum states in the open set at once: 5279405

func linearConflictsA(grid []int, size int, depth int) float32 {
	return linearConflicts(grid, size, depth) + float32(depth)
}
