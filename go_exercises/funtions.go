package main

import "fmt"

// add returns the sum of two integers.
func add(num1, num2 int) int {
	return num1 + num2
}

// divide returns the quotient and remainder of num1 / num2.
func divide(num1, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func main() {
	quo, rem := divide(10, 3)
	fmt.Println(quo, rem)
	fmt.Println(add(15, 10))
}
