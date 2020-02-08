package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, e := ioutil.ReadFile("rahu.txt")
	fmt.Println(string(dat), e)
}
