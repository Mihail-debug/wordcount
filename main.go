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
	count = len(strs)
	fmt.Println(count)
}
