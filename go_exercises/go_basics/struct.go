package main

import "fmt"

type Student struct {
	name string
	age int
}

func (stu *Student) hello(perdon string) string {
	return fmt.Sprintf("Hello, %s! I am %d years old.", stu.name, stu.age)
}

func main() {
	stu := &Student{name: "Alice"}
	stu2 := new(Student)
	stu2.name = "Bob"
	stu2.age = 20
	greeting := stu.hello("Bob")
	greeting2 := stu2.hello("Alice")
	fmt.Println(greeting2)
	fmt.Println(greeting)
}
