package heuristics

func tScore(puzzle []int, nb int, x int, y int, size int) float32 {
	nb1 := puzzle[get1d(x, y, size)]
	if nb == nb1 {
		return 0
	}
	return 1
}

// Tiles out-of place
func toop(grid []int, size int, depth int) float32 {
	var score float32
	puzzle, _ := makeGoal(size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			score += tScore(puzzle, grid[get1d(x, y, size)], x, y, size)
		}
	}
	return score
}

func toopA(grid []int, size int, depth int) float32 {
	return toop(grid, size, depth) + float32(depth)
}
