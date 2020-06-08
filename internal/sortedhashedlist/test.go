package sortedhashedlist

import "fmt"

func testMain() {
	fmt.Println("Insertion order:")
	tmp := New()
	for i := 5; i < 15; i++ {
		tmp.Insert(fmt.Sprintf("%d", i%10), float32(i%10), float32(i%10))
		fmt.Print(i % 10)
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

	tmp.Insert(fmt.Sprintf("%d", 4), float32(4), float32(4))
	fmt.Println("After add 4 (already seen):", status)
	fmt.Println(tmp)

	tmp.Insert("Q", float32(4), float32(4))
	fmt.Println("After add Q with 4 value:", status)
	fmt.Println(tmp)
}
