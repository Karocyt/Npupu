package main

import "fmt"

func main() {
	fmt.Println("Insertion order:")
	tmp := New()
	for i := 5; i < 15; i++ {
		fmt.Print(i % 10)
		tmp.insert(string(i%10), float32(i%10), float32(i%10))
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Sorted:")
	fmt.Println(tmp)
	fmt.Println()

	tmp.delete(string(4))
	fmt.Println("After delete 4:")
	fmt.Println(tmp)
}
