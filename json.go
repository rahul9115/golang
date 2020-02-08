package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
}

func main() {
	p1 := Person{name: "rahul"}

	barr, err := json.Marshal(p1)
	fmt.Println(json.Unmarshal([]byte(barr), &p1), err)
	fmt.Println(p1)

}
