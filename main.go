package main

import (
	"flag"
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
	flag.Usage()
	os.Exit(1)
}

func parseCmd() (string, map[string]solver.ScoreFn, bool) {
	var filename string
	var aStar, compare, uniform, display bool
	heuristic := 2

	hUsage := "Available heuristics:\n"
	for i, h := range heuristics.Functions {
		if i != 0 {
			hUsage += fmt.Sprintf("\t%d: %s\n", i, h.Name)
		}
	}

	flag.StringVar(&filename, "f", "", "filename of your input file")
	flag.IntVar(&heuristic, "h", 1, hUsage)
	flag.BoolVar(&aStar, "s", false, "uses A* algorithm to find the shortest path")
	flag.BoolVar(&compare, "vs", false, "compare greedy search and Astar performance")
	flag.BoolVar(&uniform, "ref", false, "adds uniform-cost search for reference")
	flag.BoolVar(&display, "display", false, "force print of full solution in any case")

	flag.Parse()

	if heuristic < 1 || heuristic >= len(heuristics.Functions) || flag.NArg() != 0 {
		flag.Usage()
		os.Exit(1)
	}

	heuristicsMap := map[string]solver.ScoreFn{}
	if compare || !aStar {
		heuristicsMap["Greedy "+heuristics.Functions[heuristic].Name] = heuristics.Functions[heuristic].Fn
	}
	if compare || aStar {
		heuristicsMap["A* "+heuristics.Functions[heuristic].Name] = heuristics.Functions[heuristic].Astar
	}
	if uniform {
		heuristicsMap[heuristics.Functions[0].Name] = heuristics.Functions[0].Fn
	}

	return filename, heuristicsMap, display
}

func main() {
	filename, heuristicsMap, display := parseCmd()
	tmp, size, e := parser.Parse(filename)
	printError(e)
	solvers := make([]*solver.Solver, 0, 2)
	solver.Init(size)
	for name, fn := range heuristicsMap {
		s := solver.New(tmp, size, fn, name)
		solvers = append(solvers, s)
		go s.Solve()
	}

	if len(solvers) == 1 {
		display = true
	}
	for i := range solvers {
		res, ok := <-solvers[i].Solution
		stats := <-solvers[i].Stats
		solvers[i].PrintRes(res, ok, stats, display)
	}
	os.Exit(0)
}
