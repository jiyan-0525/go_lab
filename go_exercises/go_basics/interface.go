package main

import "fmt"

type person interface {
	getName () string
}

type student struct {
	name string
	age int
}

func (s *student) getName() string {
	return s.name
}

type worker struct {
	name string
	geneder string
}

func (w *worker) getName() string {
	return w.name
}

func main() {
	var p person = &student{name: "Alice", age: 20}
	stu := p.(*student)
	fmt.Println(stu.age)
	fmt.Println(p.getName())
	m := make(map[string]interface{})
	m["name"] = "Bob"
	m["age"] = 30
	m["scores"] = [3]int{90, 80, 70}
	fmt.Println(m)
}
