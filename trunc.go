package main

import "fmt"

func main() {
	var number float64
	fmt.Println("Enter a floating point number :")
	fmt.Scan(&number)
	fmt.Printf("%d", int64(number))
}
