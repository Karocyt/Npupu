package heuristics

import (
	"math"
)

func eScore(nb int, x int, y int, size int) float32 {
	tmp := nbPos[nb]
	x1 := tmp[0]
	y1 := tmp[1]

	dx := float64(abs(x1 - x))
	dy := float64(abs(y1 - y))

	return float32(math.Sqrt(dx*dx + dy*dy))
}

// euclidean distance function: other basic heuristic
func euclidean(grid []int, size int, depth int) float32 {
	var score float32
	_, nbPos = makeGoal(size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			score += eScore(grid[get1d(x, y, size)], x, y, size)
		}
	}
	return score
}

func euclideanA(grid []int, size int, depth int) float32 {
	return euclidean(grid, size, depth) + float32(depth)
}
