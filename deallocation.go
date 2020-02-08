package main

import "fmt"

func foo() *int {
	x := 2
	return &x
}
func main() {
	var y *int
	y = foo()
	fmt.Println(y)
}
