package parser

import (
	"bufio"
	"errors"
	"fmt"
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

func read(filename string) (pupu []int, size int, e error) {
	fileInfo, e := os.Stat(filename)
	if e != nil {
		return
	}
	if fileInfo.IsDir() {
		e = errors.New("filename is a directory: " + filename)
		return
	}
	file, e := os.OpenFile(filename, os.O_RDONLY, 444)
	if e != nil {
		e = errors.New("failed opening file: " + filename)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		tmp := strings.TrimSpace(scanner.Text())
		if len(tmp) > 0 && tmp[0] != '#' {
			tmp := strings.Split(tmp, "#")
			size, e = strconv.Atoi(tmp[0])
			if e != nil {
				e = errors.New("bad size number")
				return
			}
			break
		}
	}
	//fmt.Println(size)
	pupu = make([]int, size*size)
	var x int
	for scanner.Scan() {
		tmp := strings.TrimSpace(scanner.Text())

		if tmp == "" || tmp[0] == '#' {
			continue
		}
		idx := len(tmp)
		for i, c := range tmp {
			if c == '#' {
				idx = i
			}
		}
		tmp = tmp[0:idx]

		tmpTab := strings.Split(tmp, " ")
		tab := make([]string, 0, size)
		for _, elem := range tmpTab {
			if elem != "" {
				tab = append(tab, elem)
			}
		}

		for i := 0; i < size; i++ {
			if len(tab) != size {
				fmt.Println(tab, len(tab))
				e = errors.New("bad size line")
				return
			}
			if len(pupu) > x*size+i {
				pupu[x*size+i], e = strconv.Atoi(tab[i])
			}
			if e != nil {
				return
			}
		}
		x++
	}
	file.Close()
	e = checkPupu(pupu, size*size)
	if e != nil {
		return
	}
	return
}

// Parse function: Only exported function
func Parse(filename string) (pupu []int, size int, e error) {
	if filename != "" {
		pupu, size, e = read(filename)
	} else {
		pupu, size = pupu_rand()
	}

	if e == nil && !checkSolvy(pupu, size) {
		fmt.Println("Pupu is not solvable :3")
		os.Exit(0)
	}
	if e != nil {
		return
	}
	return
}
