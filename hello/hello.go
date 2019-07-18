package hello

import "fmt"

// GetHello get hello
func GetHello() {
	str := "ja" + "pan"
	fmt.Println(str)
	fmt.Printf("test hello\n")
}

// Sum func
func Sum(a int, b int) int {
	return a + b
}
