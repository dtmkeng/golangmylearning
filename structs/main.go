package main

import "fmt"

// Student struct
type Student struct {
	name  string
	age   int
	grade []string
}

func main() {
	fmt.Println("test struct")
	s1 := Student{name: "keng", age: 10, grade: []string{"A", "B", "B+"}}
	// fmt.Println(s1.name)
	printdetail(s1)
	fmt.Println(s1)
}
func printdetail(students Student) {
	fmt.Println(students.name)
	printArray(students.grade)
	fmt.Println(students.age)
	students.age = 1
}
func printArray(arr []string) {
	fmt.Print("My array: ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%s ", arr[i])
	}
	fmt.Println()
}
