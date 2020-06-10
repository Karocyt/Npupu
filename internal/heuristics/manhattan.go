package heuristics

func mScore(nb int, x int, y int, size int, nbPos map[int][2]int) float32 {
	tmp := nbPos[nb]
	x1 := tmp[0]
	y1 := tmp[1]
	return float32(abs((x1 - x)) + abs((y1 - y)))
}

func manhattan(grid []int, size int, depth int) float32 {
	var score float32
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			score += mScore(grid[Get1d(x, y, size)], x, y, size, finalPos)
		}
	}
	return score
}

func manhattanA(grid []int, size int, depth int) float32 {
	return manhattan(grid, size, depth) + float32(depth)
}
