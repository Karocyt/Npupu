package main

import "fmt"

func print_case(v[] int, s int) {

	for i := 0; i < s; i++ {
		fmt.Printf("#%2d #", v[i])
	}
	fmt.Println()
	for i := 0; i < s; i++ {
		fmt.Print("#####")
	}
	fmt.Println()
}

func print_pup(pupu [][]int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print("#####")
	}
	fmt.Println()
	for i := 0 ; i < size; i++ {
			print_case(pupu[i], size)
		}

}
