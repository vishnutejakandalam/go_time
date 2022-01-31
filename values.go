package main

import "fmt"

const s string = "THIS IS A CONSTANT"

func main() {
	var a int = 10
	var b = 20
	var c bool = true
	if c {
		fmt.Print(a+b, "\n")
	}
	fmt.Printf("the constant is " + s + "\n")
	fmt.Println("Thisis a " + "string")
	// fmt.Println(1 + 12 + 23)
}
