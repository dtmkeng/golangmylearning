package main

import "fmt"

// Student sturt
type Student struct {
	name string
	age  int
}

func main() {
	a := make(map[string]int)
	b := make(map[string]*Student)
	a["1"] = 1
	a["age"] = 15
	fmt.Println(a["1"])
	fmt.Println("hell")
	s := Student{
		name: "kemg",
		age:  15,
	}
	b["1"] = &s
	fmt.Println(b["1"].age)
	s.setName("keng2")
	fmt.Println(s.getName())
}

// get method
func (s Student) getName() string {
	return s.name
}

// set method
func (s *Student) setName(name string) {
	s.name = name
}
