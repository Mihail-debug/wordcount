package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int
	var input string
	input = os.Args[1]
	strs := strings.Split(input, " ")
	for _, el := range strs {
		if el != "" {
			count += 1
		}
	}
	fmt.Println(count)
}
