package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var numbers = make([]int, 0, 3)
	var s string

	for {
		fmt.Println("Enter integer or X: ")
		fmt.Scan(&s)

		s = strings.ToLower(s)
		if s == "x" {
			return
		}

		i, _ := strconv.Atoi(s)
		numbers = append(numbers, i)
		sort.Ints(numbers)

		fmt.Println(numbers)
	}
}
