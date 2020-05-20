package parser

import (
	"bufio"
	"errors"
	"fmt"
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
		if tmp[0] == '#' {continue}
		if tmp == "" {continue}
		if len(tmp) > 0 && tmp[0] != '#' {
			tmp := strings.Split(tmp, "#")
			if tmp[0][0] == '#' {continue}
			fmt.Println(tmp)
			tmp = strings.Split(tmp[0], " ")
			for i := 0; i < size; i++ {
				if len(tmp) != size {
					log.Fatalf("bad file")
					os.Exit(-1)
				}
				pupu[x*size+i], _ = strconv.Atoi(tmp[i])
			}
		}
		x++
	}
	file.Close()

	if !check_pupu(pupu, size  * size) {
		log.Fatalf("bad file")
		os.Exit(-1)
	}
	return pupu, size
}




// Parse function: Only exported function
func Parse(heuristicsCount int) (pupu []int, size int, heuristics []int, e error) {
	pupu, size = read()
	heuristics = []int{}
	for i := 2; i < len(os.Args); i++ {
		var h int
		h, e = strconv.Atoi(os.Args[i])
		if e == nil && (h >= heuristicsCount || h < 0) {
			e = errors.New("Invalid heuristic")
		}
		if e != nil {
			return // pupu, size, heuristics, e
		}
		heuristics = append(heuristics, h)

	}
	if len(heuristics) == 0 {
		e = errors.New("Please provide a heuristic")
	}
	return
}
