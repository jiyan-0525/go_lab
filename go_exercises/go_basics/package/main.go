package main

import (
	"fmt"
	"example/calc"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())  // Ahoy, world!
	fmt.Println(calc.Add(2, 3))              // 5
}
