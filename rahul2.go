package main

import "fmt"

func main() {
	var a [2]int = [2]int{1, 2}
	slic := a[0:1]
	slic = append(slic, 1)
	fmt.Println(slic)
}
