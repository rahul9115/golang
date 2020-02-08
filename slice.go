package main

import (
	"fmt"
	"sort"
)

func main() {
	i := 1
	sli := make([]int, 3)
	value := false
	var n int
	var cont string
	for value != true {
		fmt.Println("Enter an integer")
		fmt.Scan(&n)
		i++
		sli = append(sli, n)
		sort.Ints(sli)
		fmt.Println("Do you want to continue enter X to stop or any key to continue")
		fmt.Scan(&cont)
		if cont == "X" {
			value = true
		}
		fmt.Println("The slice is ", sli)
	}
}
