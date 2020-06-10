package heuristics

func abs(nb int) int {
	if nb < 0 {
		return -nb
	}
	return nb
}

/*
func GetNInv (nb int, size int) int {
	x, y := get2d(nb, size)
	return
}
*/

func Get1d(x int, y int, size int) int {
	return x*size + y
}

func Get2d(nb int, size int) (int, int) {
	var x, y int
	x = nb / size
	y = nb % size
	return x, y
}
