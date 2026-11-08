package main

import "fmt"

// func main() {
// 	str := "Golang"
// 	var p *string = &str
// 	*p = "Hello"
// 	fmt.Println(str)
// }

func add(num int) {
	num += 1
}

func realAdd(num *int) {
	*num += 1
}

func main() {
	num := 100
	add(num)
	fmt.Println(num)  // 100 — num is unchanged

	realAdd(&num)
	fmt.Println(num)  // 101 — passed by pointer, num was modified
}