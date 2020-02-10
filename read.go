package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Names struct {
	fname string
	lname string
}

func main() {

	fname := " "
	lname := " "
	sli := []Names{}

	f, err := os.Open("name.txt")
	fmt.Println(err)

	rawBytes, err := ioutil.ReadAll(f)

	fmt.Println(err)
	lines := strings.Split(string(rawBytes), "\n")
	for i, line := range lines {
		k := 0
		fmt.Printf("%d", i)
		for j, c := range line {
			fmt.Printf("%d", j)
			if string(c) == " " {
				k = 1
				continue
			}
			fmt.Println(k)
			if string(c) != " " && k != 1 {
				fname = fname + string(c)
			}
			if string(c) != " " && k == 1 {

				lname = lname + string(c)

			}

		}
		// names := {
		// 	"first_name": a.fname,
		// 	"last_name":  a.lname}
		// sli = append(sli, names)
		n1 := Names{fname, lname}
		sli = append(sli, n1)
		fname = " "
		lname = " "

	}
	for i, line := range sli {
		fmt.Println(i, line)
	}

}
