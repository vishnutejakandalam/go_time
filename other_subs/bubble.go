package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

func main() {
	
	inp_slice := make([]int, 0, 10)

	for i := 1; i <= 10; i++ {
		fmt.Printf("Enter integer %d out of 10: ", i)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		datapoint, _  := strconv.ParseInt(scanner.Text(), 0, 64)
		inp_slice = append(inp_slice, int(datapoint))
	}

	BubbleSort(inp_slice)
	
}

func BubbleSort(inp []int) {

	for i := range inp {
		for j := 0; j < len(inp)-i-1; j++ {
			if inp[j] > inp[j+1] {
				Swap(&inp[j], &inp[j+1])
			}
		}
	}
    for element := range inp {
		fmt.Printf("%d ", inp[element])
	}
	
}

func Swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
 }
