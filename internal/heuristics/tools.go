package heuristics

func abs(nb int) int {
	if 	nb < 0 {
		return -nb
	} else {
		return nb
	}
}

func get_1d(x int, y int, size int) int {
	return x * size + y
}

func get_2d(nb int, size int) (int, int) {
	var x, y int
	x = nb / size
	y = nb % size
	return x, y
}