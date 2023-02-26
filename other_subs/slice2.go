package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ints := make([]int, 3)
	fmt.Println("Enter integers. Enter X to exit.")
	var input string

	i := 0

	for {

		fmt.Scanln(&input)

		if strings.ToLower(input) == "x" {
			break
		}

		num, _ := strconv.Atoi(input)

		if i < 3 {
			ints[0] = num
		} else {
			ints = append(ints, num)
		}

		sort.Ints(ints)
		fmt.Println(ints)

		i++

	}

}
