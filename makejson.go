package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	reader := bufio.newReader(os.Stdin)
	var name, address string
	var p1 map[string]string
	fmt.Println("Enter your first name")
	name, _ := reader.readString("\n")
	fmt.Println("Enter your address")
	address, _ := reader.readString("\n")
	idmap := map[string]string{
		"name":    name,
		"address": address}
	object, err := json.Marshal(idmap)
	fmt.Println(object, err)
	json.Unmarshal([]byte(object), &p1)
	fmt.Println(p1)
}
