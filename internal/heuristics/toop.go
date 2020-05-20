package heuristics

func tScore(grid []int, nb int, x int, y int, size int) float32 {
	tmp := nb_pos[nb]
	x1 := tmp[0]
	y1 := tmp[1]
	nb1 := grid[get1d(x1, y1, size)]
	if nb == nb1 {
		return 1
	} else {
		return 0
	}
	return 0
}

// Tiles out-of place
func toop(grid []int, size int) float32 {
	var score float32
	_, nb_pos = makeGoal(size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			score += tScore(grid, grid[get1d(x, y, size)], x, y, size)
		}
	}
	return score
}
