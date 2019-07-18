package main
import "fmt"
// Student sturt
type Student struct {
	name string
	age int
}
func main() {
	fmt.Println("hell")
	s := Student{
		name: "kemg",
		age: 15,
	}	
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
