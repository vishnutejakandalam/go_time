package main

import (
	"fmt"
	"time"
)

func sample_1(x *int) {
	time.Sleep(7 * time.Second)
	fmt.Println(*x)
}

func sample_2(x *int) {
	time.Sleep(5 * time.Second)
	*x = *x + 1
	fmt.Println(*x)
}

func main() {
	x := 1
	fmt.Println("start")
	go sample_1(&x)
	go sample_2(&x)
	time.Sleep(8 * time.Second)
	fmt.Println("expexted order value: 1, 2")
	fmt.Println("but function sample_2 executed more fast that sample_1")
	fmt.Println("eventualy will return 2 2, these is race condition - var x war rewrite fuction sample_2 before sample_1 print it")
}
