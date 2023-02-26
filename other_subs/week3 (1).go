package main

import(
	"fmt"
	"strconv"
)

func main(){
	numbers := make([]int,0)
	var count int = 0
	fmt.Println("Enter at least 12 integers to sort and X when you are done")
	for {
		fmt.Println("Enter Integer:")
		var in string
		fmt.Scan(&in)
		if in == "X" && count<=12{ 
			break 
		}else if in == "X"{
			fmt.Println("Continue until you entered at least 12 integers")
		}else{ 
			n , _ := strconv.Atoi(in)	
			numbers = append(numbers,n)		
		}
	}
	size := int(len(numbers)/4)
	array1 := numbers[0:size]
	array2 := numbers[size:2*size]
	array3 := numbers[2*size:3*size]
	array4 := numbers[3*size:len(numbers)]
	c := make(chan []int,4)
	go BubbleSort(array1,c)
	go BubbleSort(array2,c)
	go BubbleSort(array3,c)
	go BubbleSort(array4,c)
	array1 = <- c
	array2 = <- c
	array3 = <- c
	array4 = <- c
	final := merge(merge(array1,array2),merge(array3,array4))
	fmt.Println(final)

}

func BubbleSort(seq []int,c chan []int) {
	for i:= 0 ; i < len(seq) - 1 ; i++ {
		for j:=0; j < len(seq)-i-1; j++ {
			if seq[j] > seq[j+1]{
				Swap(seq,j)
			}
		}
	}
	c <- seq
}

func Swap(seq []int, i int) {
	temp := seq[i]
	seq[i] = seq[i+1]
	seq[i+1] = temp
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}