package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("This is loop of i", i)
	}
	var temp = 1
	for temp <= 10 {
		fmt.Print("While for for ", temp, "\n")
		temp++
	}

}
