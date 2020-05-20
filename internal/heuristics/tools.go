package heuristics

func abs(nb int) int {
	if nb < 0 {
		return -nb
	} else {
		return nb
	}
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

func makeGoal(s int) []int {
	puzzle := make([]int, s*s)
	cur := 1
	x := 0
	ix := 1
	y := 0
	iy := 0
	for cur < s*s {
		puzzle[x+y*s] = cur

		cur++

		if x+ix == s || x+ix < 0 || (ix != 0 && puzzle[x+ix+y*s] != 0) {
			iy = ix
			ix = 0
		} else if y+iy == s || y+iy < 0 || (iy != 0 && puzzle[x+(y+iy)*s] != 0) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
	}
	puzzle[x+y*s] = 0

	return puzzle
}
