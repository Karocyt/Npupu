package main

import (
	"os"

	"github.com/Karocyt/Npupu/internal/parser"
)

func mainfunc() int {
	tmp, size := parser.Parse()
	print_pup(tmp, size)
	return (0)
}

func main() {
	os.Exit(mainfunc())
}
