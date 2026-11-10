package main

import "fmt"

// func add(num1 int, num2 int) int {
//     return num1 + num2
// }

func add(num1 int, num2 int) (ans int) {
	ans = num1 + num2
	return
}
func divide(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}


func main() {
	quo, rem:= divide(10, 3)
	fmt.Println(quo, rem)
    fmt.Println(add(15, 10))
}