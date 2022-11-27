package main

import (
	"fmt"
	"sort"
	"strconv"
	"syscall"
)

func main() {
	var slice = make([]int, 0, 3)
	for {
		var val string
		fmt.Println("Enter a value to enter into the slice or X to exit: ")
		fmt.Scanf("%s", &val)
		if val == "X" {
			fmt.Println("Exiting the program..")
			syscall.Exit(0)
		}
		int_val, _ := strconv.ParseInt(val, 0, 16)
		slice = append(slice, int(int_val))
		sort.Sort(sort.IntSlice(slice))
		fmt.Println("After sorting the slice is: ", slice)
	}
}
