package main

import (
	"fmt"
	"sync"
)

func double(x *int, wg *sync.WaitGroup) {
	fmt.Printf("Doubling...")
	defer wg.Done()
	*x = *x * 2
}

func increaseByOne(x *int, wg *sync.WaitGroup) {
	fmt.Printf("add one. ...")
	defer wg.Done()
	*x = *x + 1
}

func main() {

	x := 0
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 5; i++ {
		// Race condition occurs due to the access of the same shared x variable
		// and to the fact that we can not determinate the order of execution
		// some sync in needed to get the desired behavior
		// otherwise each time run the program, the result may be different
		go increaseByOne(&x, &wg)
		go double(&x, &wg)
	}
	wg.Wait()
	fmt.Println("result", x)
}
