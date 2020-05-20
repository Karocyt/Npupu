package main

import (
	"fmt"
	"os"

	"github.com/Karocyt/Npupu/internal/heuristics"
	"github.com/Karocyt/Npupu/internal/parser"
	"github.com/Karocyt/Npupu/internal/solver"
)

func printSolution(s solver.Solver, name string) {
	fmt.Printf("Solution using %s:\n\n", name)
	for _, step := range s.Solution {
		fmt.Println(step)
	}
}

func mainfunc() int {
	tmp, size, h, e := parser.Parse(len(heuristics.Functions))
	if e != nil {
		fmt.Fprint(os.Stderr, e)
		return 1
	}
	for _, currH := range h {
		s := solver.New(tmp, size, heuristics.Functions[currH].Fn)
		e = s.Solve()
		if e != nil {
			fmt.Fprint(os.Stderr, e)
			return 1
		}
		printSolution(s, heuristics.Functions[currH].Name)
	}

	return (0)
}

func main() {
	os.Exit(mainfunc())
}
