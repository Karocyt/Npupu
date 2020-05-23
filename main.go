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

func validateArgs() {
	if len(os.Args) < 2 {
		printError(errors.New("Please provide a file to open"))
	}
}

func mainfunc() int {
	validateArgs()
	tmp, size, h, e := parser.Parse(len(heuristics.Functions))
	printError(e)
	for _, currH := range h {
		s := solver.New(tmp, size, heuristics.Functions[currH].Fn)
		e = s.Solve()
		printError(e)
		s.PrintRes(heuristics.Functions[currH].Name)
	}

	return (0)
}

func main() {
	os.Exit(mainfunc())
}
