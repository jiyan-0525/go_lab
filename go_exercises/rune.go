package main

import (
	"fmt"
	"reflect"
)
func main() {
	str1 := "Golang"
	str2 := "Go语言"
	runearr := []rune(str2)
	fmt.Println(reflect.TypeOf(runearr[2]).Kind())
	fmt.Println(runearr[2], string(str1[2]))
	fmt.Printf("%d, %c\n", runearr[2], runearr[2])
	fmt.Println("len(str2):", len(runearr))
}