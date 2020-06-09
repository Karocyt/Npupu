package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Karocyt/Npupu/internal/heuristics"
	"github.com/Karocyt/Npupu/internal/parser"
	"github.com/Karocyt/Npupu/internal/solver"
)

func printError(e error) {
	if e == nil {
		return
	}
	fmt.Print("ERROR: ")
	fmt.Fprintln(os.Stderr, e)
	fmt.Printf("usage: %s filename [0 to %d heuristic]\n\n", os.Args[0], len(heuristics.Functions)-1)
	fmt.Println("Available heuristics:")
	for i, h := range heuristics.Functions {
		fmt.Printf("\t%d: %s\n", i, h.Name)
	}
	os.Exit(1)
}

func validateArgs() (e error) {
	if len(os.Args) < 2 {
		printError(errors.New("Please provide a file to open"))
	} else if len(os.Args) > 4 {
		printError(errors.New("Too Many arguments"))
	}
	return
}

func main() {
	validateArgs()
	tmp, size, h, e := parser.Parse(len(heuristics.Functions))
	printError(e)
	solvers := make([]*solver.Solver, 0, 2)
	solver.Init(size)
	for _, currH := range h {
		solvers = append(solvers, solver.New(tmp, size, heuristics.Functions[currH].Fn))
		go solvers[len(solvers)-1].Solve()
	}

	for i := range solvers {
		res, ok := <-solvers[i].Solution
		stats := <-solvers[i].Stats
		solvers[i].PrintRes(heuristics.Functions[h[i]].Name, res, ok, stats)
	}
	os.Exit(0)
}
