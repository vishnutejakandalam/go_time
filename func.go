package main

import "fmt"

func hello(a int, b int) int {
	fmt.Println("Called function");
	return a+b;
}

func swap(a *int, b *int){
	var temp int;
	temp = *a;
	*a = *b;
	*b = temp;
}

func main() {
	fmt.Println("Calling the fuction:");
	var c = hello(12,13);
	fmt.Println("The sum is ",c);
	var a int = 10;
	var b int = 20;

	fmt.Println("The values before swapping are: ", a, b)
	swap(&a, &b)
	fmt.Println("The values before swapping are: ", a, b)


}

