package main

import "fmt"

func main() {

type Gender int8
const (
	MALE Gender = 1
	FEMALE Gender = 2
)

gender := MALE

switch gender {
case MALE:
	fmt.Println("Male")
case FEMALE:
	fmt.Println("Female")
default:
	fmt.Println("Unknown")
}
}