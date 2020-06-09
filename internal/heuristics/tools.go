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
var Format string

func Get1d(x int, y int, size int) int {
	return x*size + y
}

func Get2d(nb int, size int) (int, int) {
	var x, y int
	x = nb / size
	y = nb % size
	return x, y
}

func MakeGoal(s int) ([]int, map[int][2]int) {
	nbPos := make(map[int][2]int)
	puzzle := make([]int, s*s)
	cur := 1
	x := 0
	ix := 1
	y := 0
	iy := 0
	if Format == "snail" {
		for cur < s*s {
			puzzle[x+y*s] = cur
			nbPos[cur] = [2]int{y, x}
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
		nbPos[0] = [2]int{y, x}
		puzzle[x+y*s] = 0
	}
	if Format == "classic" {
		for cur < s*s {
			if x == s {
				y++
				x = 0
			}
			puzzle[x+y*s] = cur
			nbPos[cur] = [2]int{y, x}
			cur++

			x++
		}
		nbPos[0] = [2]int{y, x}
		puzzle[x+y*s] = 0
	}
	return puzzle, nbPos

}
