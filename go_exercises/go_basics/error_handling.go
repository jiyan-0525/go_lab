package main

import (
	"fmt"
	// "os"
	"errors"
)

// func main() {
// 	_, err := os.Open("nonexistent_file.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

func hello(name string) error {
	// name = "Alice"
	if len(name) == 0  {
		return errors.New("name cannot be empty")
	}
	fmt.Printf("Hello, %s!\n", name)
	return nil
}

func getdigit(n int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			ret = -1 // Return a default value in case of panic
		}
	}()
	arr := [3]int{1, 2, 3}
	// This will cause a panic if n is out of bounds
	return arr[n]
}

func main () {
	fmt.Println("digits:", getdigit(3)) // This will cause a panic and be recovered
	fmt.Println("finished")
}

// func main() {
// 	if err := hello(""); err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }
