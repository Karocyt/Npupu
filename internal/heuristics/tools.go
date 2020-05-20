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

func make_goal(s int ) []int {
	ts := s * s
	puzzle := make([]int, ts)
	cur := 1
	x := 0
	ix := 1
	y := 0
	iy := 0
	for {
		puzzle[x+y*s] = cur

		if cur == 0 {
			break
		}
		cur += 1

		if x+ix == s || x+ix < 0 || (ix != 0 && puzzle[x+ix+y*s] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == s || y+iy < 0 || (iy != 0 && puzzle[x+(y+iy)*s] != -1) {
			ix = -iy
			iy = 0
			x += ix
			y += iy
		}
		if cur == s*s {
			cur = 0
		}
	}
	return puzzle
}
