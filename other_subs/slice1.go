package main
  
import (
    "fmt"
    "sort"
)

func main() {
	
	my_slice := []int{12, 45, 67}
	
	for {
		fmt.Print("\nWant to Add Data In Slice (Enter Anything to continue or 'X' to Break): ")
		var str string
		fmt.Scanln(&str)
		
		if str == "X"{
			break
		}
		fmt.Print("Enter an Integer: ")
		var num int
		fmt.Scanln(&num)
		
		res := append(my_slice, num)
		my_slice = res
		
		fmt.Print("New Sorted Slice after Adding an New Data: \n")
		
		sort.Slice(my_slice, func(i, j int) bool {
			return my_slice[i] < my_slice[j]
		})
		
		fmt.Print(my_slice)
		
    }
}