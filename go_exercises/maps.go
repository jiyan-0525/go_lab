package main

import "fmt"

func main() {
	mi := make(map[string]int)
	mi["one"] = 1
	mi["two"] = 2
	mi["three"] = 3
	
	m2 := map[string]string {
		"one": "uno",
		"two": "dos",
		"three": "tres",
	}
	fmt.Println(mi)
	fmt.Println(mi["two"])
	fmt.Println(mi["four"])
	fmt.Println(m2)
	fmt.Println(m2["two"])
	fmt.Println(m2["four"])
}