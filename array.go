package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 4, 5, 6}
	slic := a[1:5]
	fmt.Println(len(slic), cap(slic))

}
