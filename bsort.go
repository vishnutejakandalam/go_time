package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	var temp int
	// fmt.Println("Sorting the nubmers..")
	temp = *a
	*a = *b
	*b = temp
}

func bsort(list []int) {
	// fmt.Println("Entered the sorting function: ")
	n := len(list)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if list[i] > list[j] {
				swap(&list[i], &list[j])
			}
		}
	}
}

func main() {
	list := make([]int, 0)
	// list = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Enter 10 numbers for the list.")
	for i := 0; i < 10; i++ {
		val := 0
		fmt.Scanf("%d", &val)
		list = append(list, val)
	}
	fmt.Print("The values are ")
	fmt.Println(list)
	bsort(list)
	fmt.Print("After sorting the values are ")
	fmt.Println(list)

}
