package heuristics

import (
	"math"
)

func eScore(nb int, x int, y int, size int, nbPos map[int][2]int) float32 {
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
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			val := grid[Get1d(x, y, size)]
			if val != 0 {
				score += eScore(val, x, y, size, finalPos)
			}
		}
	}
	return score
}

func euclideanA(grid []int, size int, depth int) float32 {
	return euclidean(grid, size, depth) + float32(depth)
}
