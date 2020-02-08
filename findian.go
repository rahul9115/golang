package main

import "fmt"

func main() {
	var str string
	var i int = 0
	var value bool = false
	fmt.Println("Enter a string")
	fmt.Scan(&str)
	for ; i < len(str); i++ {

		if i == 0 {
			if string(str[i]) != "i" {
				break
			}
		}
		if string(str[i]) == "a" && i != 0 && i != (len(str)-1) {
			value = true
		}

		if i == (len(str) - 1) {
			if string(str[i]) != "n" {
				break
			}
		}

	}
	fmt.Println(i)
	if i == len(str) && value == true {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
