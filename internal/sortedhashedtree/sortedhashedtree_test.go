package sortedhashedtree

import (
	"fmt"
	"testing"
)

// TestUno is a test function to use with the go test utility
func TestUno(*testing.T) {
	fmt.Println("#### MIX ########################################")
	fmt.Println("Insertion order:")
	tmp := New()
	for i := 5; i < 15; i++ {
		tmp.Insert(fmt.Sprintf("%d", i*2%20), float32(i*2%20), float32(i*2%20))
		fmt.Print(i * 2 % 20)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Sorted:")
	fmt.Println(tmp)
	fmt.Println()

	status := tmp.Delete(fmt.Sprintf("%d", 4))
	fmt.Println("After delete 4:", status)
	fmt.Println(tmp)

	status = tmp.Delete(fmt.Sprintf("%d", 1))
	fmt.Println("After delete 1:", status)
	fmt.Println(tmp)

	status = tmp.Delete(fmt.Sprintf("%d", 5))
	fmt.Println("After delete 5:", status)
	fmt.Println(tmp)

	// status = tmp.Delete(fmt.Sprintf("%d", 9))
	// fmt.Println("After delete 9:", status)
	// fmt.Println(tmp)

	status = tmp.Delete(fmt.Sprintf("%d", 0))
	fmt.Println("After delete 0:", status)
	fmt.Println(tmp)

	tmp.Insert(fmt.Sprintf("%d", 4), float32(4), float32(4))
	fmt.Println("After add 4 (already seen):", status)
	fmt.Println(tmp)

	tmp.Insert("Q", float32(4), float32(4))
	fmt.Println("After add Q with 4 value:", status)
	fmt.Println(tmp)

	tmp.Insert("15", float32(15), float32(15))
	fmt.Println("After add 15:", status)
	fmt.Println(tmp)

	tmp.Insert("15.5", float32(15.5), float32(15.5))
	fmt.Println("After add 15.5:", status)
	fmt.Println(tmp)

	tmp.Insert("14.9", float32(14.9), float32(14.9))
	fmt.Println("After add 14.9:", status)
	fmt.Println(tmp)

	status = tmp.Delete("16")
	fmt.Println("After delete 16:", status)
	fmt.Println(tmp)

	status = tmp.Delete("8")
	fmt.Println("After delete 8:", status)
	fmt.Println(tmp)
}

// TestDos is a test function to use with the go test utility
func TestDos(*testing.T) {
	fmt.Println("#### delete right #################################################")
	tmp := New()
	tmp.Insert("1", 1, 1.0)
	tmp.Insert("0", 0, 0.0)
	tmp.Insert("2", 2, 2.0)
	fmt.Println("Sorted:")
	fmt.Println(tmp)

	status := tmp.Delete("2")
	fmt.Println("After delete 2:", status)
	fmt.Println(tmp)
}

// TestTres is a test function to use with the go test utility
func TestTres(*testing.T) {
	fmt.Println("#### delete left #################################################")
	tmp := New()
	tmp.Insert("1", 1, 1.0)
	tmp.Insert("0", 0, 0.0)
	tmp.Insert("2", 2, 2.0)
	fmt.Println("Sorted:")
	fmt.Println(tmp)

	status := tmp.Delete("0")
	fmt.Println("After delete 0:", status)
	fmt.Println(tmp)
}

// TestQuatro is a test function to use with the go test utility
func TestQuatro(*testing.T) {
	fmt.Println("#### delete header #################################################")
	tmp := New()
	tmp.Insert("1", 1, 1.0)
	tmp.Insert("0", 0, 0.0)
	tmp.Insert("2", 2, 2.0)
	fmt.Println("Sorted:")
	fmt.Println(tmp)

	status := tmp.Delete("0")
	fmt.Println("After delete 0:", status)
	fmt.Println(tmp)
}

// Test5 is a test function to use with the go test utility
func Test5(*testing.T) {
	fmt.Println("#### npupu fail case #################################################")
	tmp := New()
	tmp.Insert("18-0", 18, 18.0)
	fmt.Println(tmp)
	tmp.Insert("18-1", 18, 18.0)
	fmt.Println(tmp)
	tmp.Insert("18-2", 18, 18.0)
	fmt.Println(tmp)
	tmp.Delete("18-0")
	fmt.Println(tmp)
	tmp.Insert("20", 20, 20.0)
	fmt.Println(tmp)
	tmp.Insert("16-0", 16, 16.0)
	fmt.Println(tmp)
	tmp.Delete("18-1")
	// fmt.Println(tmp)
	// tmp.Insert("16-1", 16, 16.0)
	// fmt.Println(tmp)
	// tmp.Insert("16-2", 16, 16.0)
	// fmt.Println("Sorted:")
	// fmt.Println(tmp)

	status := tmp.Delete("0")
	fmt.Println("After delete 0:", status)
	fmt.Println(tmp)
}
