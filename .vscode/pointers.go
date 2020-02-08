package main

import "fmt"

func main() {
	var x int = 2
	var y int
	var ip *int
	ip = &x
	y = *ip
	fmt.Println(y)
	fmt.Scan(&y)
	fmt.Println(y)
}
