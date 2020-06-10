package heuristics

func abs(nb int) int {
	if nb < 0 {
		return -nb
	}
	return nb
}

func get1d(x int, y int, size int) int {
	return x*size + y
}

func get2d(nb int, size int) (int, int) {
	var x, y int
	x = nb / size
	y = nb % size
	return x, y
}
