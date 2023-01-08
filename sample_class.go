package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s *student) set_name(name string) {
	fmt.Println("The name is ", name)
	s.name = name
}

func (s *student) get_age() {
	fmt.Println("The age of student is ", s.age)
}

func main() {
	obj := student{}
	obj.age = 12
	obj.set_name("raju")
	obj.get_age()
	fmt.Println("the name is ", obj.name)
}
