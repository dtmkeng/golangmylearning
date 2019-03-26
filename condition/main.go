package main

import "fmt"

func main() {

	fmt.Println("Condition: ")
	if x := 1; x >= 0 {
		fmt.Println("x is more than 0")
	} else if x == 2 {
		fmt.Println("x == ", x)
	} else {
		fmt.Println("x is not scop")
	}
}
