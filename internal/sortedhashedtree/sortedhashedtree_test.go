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

	status = tmp.Delete(fmt.Sprintf("%d", 10))
	fmt.Println("After delete 10:", status)
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

	status = tmp.Delete("0")
	fmt.Println("After delete 0:", status)
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

// TestOLeft is a test function to use with the go test utility
func TestOLeft(*testing.T) {
	fmt.Println("#### delete node with only left child ###################################")
	tmp := New()
	tmp.Insert("1", 1, 1.0)
	fmt.Println(tmp)
	tmp.Insert("2", 2, 2.0)
	fmt.Println(tmp)
	tmp.Insert("0", 0, 0.0)
	fmt.Println(tmp)

	status := tmp.Delete("2")
	fmt.Println("After delete 2:", status)
	fmt.Println(tmp)
}

// TestORight is a test function to use with the go test utility
func TestORight(*testing.T) {
	fmt.Println("#### delete node with only right child ###################################")
	tmp := New()
	tmp.Insert("1", 18, 18.0)
	fmt.Println(tmp)
	tmp.Insert("2", 17, 19.0)
	fmt.Println(tmp)
	tmp.Insert("3", 16, 20.0)
	fmt.Println(tmp)

	status := tmp.Delete("2")
	fmt.Println("After delete 2:", status)
	fmt.Println(tmp)
}

// TestDelBoth is a test function to use with the go test utility
func TestDelBoth(*testing.T) {
	fmt.Println("#### delete node with both childs ###################################")
	tmp := New()
	tmp.Insert("1", 1, 1.0)
	fmt.Println(tmp)
	tmp.Insert("2", 2, 2.0)
	fmt.Println(tmp)
	tmp.Insert("1.5", 3, 1.5)
	fmt.Println(tmp)
	tmp.Insert("3", 3, 3.0)
	fmt.Println(tmp)

	status := tmp.Delete("2")
	fmt.Println("After delete 2:", status)
	fmt.Println(tmp)
}

// TestDelNoChild is a test function to use with the go test utility
func TestDelNoChild(*testing.T) {
	fmt.Println("#### delete node with no child ###################################")
	tmp := New()
	tmp.Insert("1", 1, 18.0)
	fmt.Println(tmp)
	tmp.Insert("2", 2, 17.0)
	fmt.Println(tmp)

	status := tmp.Delete("2")
	fmt.Println("After delete 2:", status)
	fmt.Println(tmp)
}

// TestDelHead is a test function to use with the go test utility
func TestDelHead(*testing.T) {
	fmt.Println("#### delete node with both childs ###################################")
	tmp := New()
	tmp.Insert("1", 1, 18.0)
	fmt.Println(tmp)
	tmp.Insert("2", 2, 17.0)
	fmt.Println(tmp)
	tmp.Insert("3", 3, 16.0)
	fmt.Println(tmp)
	tmp.Insert("4", 4, 17.5)
	fmt.Println(tmp)
	tmp.Insert("5", 4, 19.0)
	fmt.Println(tmp)
	tmp.Insert("5", 4, 20.0)
	fmt.Println(tmp)

	status := tmp.Delete("1")
	fmt.Println("After delete 1:", status)
	fmt.Println(tmp)
}

// TestReAdd is a test function to use with the go test utility
func TestReAdd(*testing.T) {
	fmt.Println("#### re-add deleted node ###################################")
	tmp := New()
	tmp.Insert("1", 1, 18.0)
	fmt.Println(tmp)
	tmp.Insert("2", 2, 17.0)
	fmt.Println(tmp)
	tmp.Insert("3", 3, 16.0)
	fmt.Println(tmp)
	tmp.Insert("4", 4, 17.5)
	fmt.Println(tmp)
	tmp.Insert("5", 4, 19.0)
	fmt.Println(tmp)
	tmp.Insert("5", 4, 20.0)
	fmt.Println(tmp)

	status := tmp.Delete("2")
	fmt.Println("After delete 2:", status)
	fmt.Println(tmp)

	tmp.Insert("2", 2, 17.0)
	fmt.Println(tmp)

	tmp.Insert("2", 2, 17.0)
	fmt.Println(tmp)
}

// TestZero is a test function to use with the go test utility
func TestZero(t *testing.T) {
	fmt.Println("#### Fucking Zero ########################################")
	fmt.Println("Insertion order:")
	tmp := New()
	for i := 5; i < 15; i++ {
		tmp.Insert(fmt.Sprintf("%d", i*2%20), float32(i*2%20), float32(i*2%20))
		fmt.Print(i*2%20, " ")
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

	status = tmp.Delete(fmt.Sprintf("%d", 10))
	fmt.Println("After delete 10:", status)
	fmt.Println(tmp)

	status = tmp.Delete(fmt.Sprintf("%d", 0))
	fmt.Println("After delete 0:", status)
	fmt.Println(tmp)

	status = tmp.Delete(fmt.Sprintf("%d", 0))
	fmt.Println("After delete 0:", status)
	fmt.Println(tmp)

	if status == true {
		t.Errorf("ERROR: 10 deletion doesn't update 0's parent")
	}
}

// TestRB1 is a test function to use with the go test utility
func TestRB1(t *testing.T) {
	fmt.Println("#### test insert right rebalancing ########################################")
	fmt.Println("Insertion order:")
	tmp := New()
	for i := 0; i < 10; i++ {
		tmp.Insert(fmt.Sprintf("%d", i), float32(i), float32(i))
		fmt.Println(i, " ")
		fmt.Println(tmp)
	}
	res := "(((0)1(2))3((4)5((6)7(8(9)))))"
	if tmp.String() != res {
		t.Errorf("'%s' should be '%s'", tmp.String(), res)
	}
}

// TestRB2 is a test function to use with the go test utility
func TestRB2(t *testing.T) {
	fmt.Println("#### test insert left rebalancing ########################################")
	fmt.Println("Insertion order:")
	tmp := New()
	for i := 10; i > 0; i-- {
		tmp.Insert(fmt.Sprintf("%d", i), float32(i), float32(i))
		fmt.Println(i, " ")
		fmt.Println(tmp)
	}
	res := "(((((1)2)3(4))5(6))7((8)9(10)))"
	if tmp.String() != res {
		t.Errorf("'%s' should be '%s'", tmp.String(), res)
	}
}

// TestRB3 is a test function to use with the go test utility
func TestRB3(t *testing.T) {
	fmt.Println("#### test delete rebalancing ########################################")
	fmt.Println("Insertion order:")
	tmp := New()
	for i := 0; i < 20; i++ {
		tmp.Insert(fmt.Sprintf("%d", i), float32(i), float32(i))
		fmt.Println(tmp)
	}

	for i := 10; i < 15; i++ {
		tmp.Delete(fmt.Sprintf("%d", i))
		fmt.Println(i, " ")
		fmt.Println(tmp)
	}
	res := "(((0)1(2))3((4)5((6)7(8(9)))))"
	if tmp.String() != res {
		t.Errorf("'%s' should be '%s'", tmp.String(), res)
	}
}

// TestRB4 is a test function to use with the go test utility
func TestRB4(t *testing.T) {
	fmt.Println("#### test red header ########################################")
	fmt.Println("Insertion order:")
	tmp := New()
	for i := 0; i < 20; i++ {
		tmp.Insert(fmt.Sprintf("%d", i), float32(i), float32(i))
		fmt.Println(tmp)
	}

	for i := 8; i < 20; i++ {
		tmp.Delete(fmt.Sprintf("%d", i))
		fmt.Println(i, " ")
		fmt.Println(tmp)
	}
	tmp.Delete(fmt.Sprintf("%d", 7))
	fmt.Println(7, " ")
	fmt.Println(tmp)
	res := "(((0)1(2))3((4)5((6)7(8(9)))))"
	if tmp.String() != res {
		t.Errorf("'%s' should be '%s'", tmp.String(), res)
	}
}
