package main

import (
	"fmt"
	"sync"
)

func go_sort(list []int, wg *sync.WaitGroup) {
	// fmt.Println("Entered the sorting function: ")
	n := len(list)
	fmt.Println("Soring this sub array: ", list)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if list[i] > list[j] {
				temp := list[i]
				list[i] = list[j]
				list[j] = temp
			}
		}
	}
	wg.Done()
}

func main() {
	var n int
	var wg = new(sync.WaitGroup)
	var values []int
	set_values := make([][]int, 0)
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		temp := 0
		fmt.Scanf("%d", &temp)
		values = append(values, temp)
		if i%4 == 0 {
			set_values = append(set_values, values)
			values = make([]int, 0)
		}
	}
	set_values = append(set_values, values)

	for _, list := range set_values {
		wg.Add(1)
		go go_sort(list, wg)
	}

	complete_list := make([]int, 0)
	for _, vals := range set_values {
		for _, val := range vals {
			complete_list = append(complete_list, val)
		}
	}

	wg.Add(1)
	go_sort(complete_list, wg)
	wg.Wait()

	fmt.Println("\n Final Sorted list", complete_list)
}
