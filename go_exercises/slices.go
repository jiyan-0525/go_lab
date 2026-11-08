package main

import(
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}

	slice1 := make([]float32, 0)
	slice2 := make([]float32, 3, 5)
	slice2 = append(slice2, 1, 2, 3, 4, 5)
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))
	sub1 := slice2[:3]
	sub2 := slice2[3:]
	sub3 := slice2[1:4]
	fmt.Println("sub1:", sub1)
	fmt.Println("sub2:", sub2)
	fmt.Println("sub3:", sub3)
	combine := append(sub1, sub2...)
	fmt.Println("combine:", combine)
	// fmt.Printf("slice1: %v, len: %d, cap: %d\n", slice1, len(slice1), cap(slice1))
    fmt.Printf("slice2: %v, len: %d, cap: %d\n", slice2, len(slice2), cap(slice2))
	// fmt.Println("arr:", arr)

}