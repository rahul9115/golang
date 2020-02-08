package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

func main() {
	numbers := make([]int, 3)
	scanner := bufio.NewScanner(os.Stdin)
	userPrompt := ""

	for strings.ToLower(userPrompt) != "x" {
		fmt.Println("Type a number or press X to exit: ")
		scanner.Scan()
		userPrompt = scanner.Text()

		if govalidator.IsInt(userPrompt) {
			value, _ := strconv.Atoi(userPrompt)
			numbers = append(numbers, value)
			sort.Ints(numbers)
			printSlice("numbers", numbers)
		}
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
