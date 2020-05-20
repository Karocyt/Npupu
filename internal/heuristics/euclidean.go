package heuristics

import (
	"math"
)

func e_score(nb int, x int, y int, size int) float32 {
	x1, y1 := get_2d(nb, size)

	dx := float64(abs(x1 - x))
	dy := float64(abs(y1 - y))

	return float32(math.Sqrt(dx * dx + dy * dy))
}

// euclidean distance function: other basic heuristic
func euclidean(grid []int, size int) float32 {
	var score float32
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			score += e_score(grid[get_1d(x , y, size)], x, y, size)
		}
	}
	return score
}
