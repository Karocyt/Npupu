package parser

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var size int

func isValid(str string) bool {
	for c := range str {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func read() ([]int, int) {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var pupu []int
	var size int

	for scanner.Scan() {
		tmp := strings.TrimSpace(scanner.Text())
		if len(tmp) > 0 && tmp[0] != '#' {
			tmp := strings.Split(tmp, "#")
			size, _ = strconv.Atoi(tmp[0])
			break
		}
	}
	//fmt.Println(size)
	pupu = make([]int, size*size)
	var x int
	for scanner.Scan() {
		tmp := strings.TrimSpace(scanner.Text())
		if len(tmp) > 0 && tmp[0] != '#' {
			tmp := strings.Split(tmp, "#")
			tmp = strings.Split(tmp[0], " ")
			for i := 0; i < size; i++ {
				pupu[x*size+i], _ = strconv.Atoi(tmp[i])
			}
		}
		x++
	}
	file.Close()

	return pupu, size
}

// Parse function: Only exported function
func Parse(heuristicsCount int) (pupu []int, size int, heuristics []int, e error) {

	pupu, size = read()
	heuristics = []int{0}
	return
}
