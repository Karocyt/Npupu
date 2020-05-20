package heuristics

func m_score(nb int, x int, y int, size int) float32 {
	x1, y1 := get_2d(nb, size)
	return float32(abs((x1 - x)) + abs((y1 - y)))
}

func manhattan(grid []int, size int) float32 {
	var score float32
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			score += m_score(grid[get_1d(x , y, size)], x, y, size)
		}
	}
	return score
}
