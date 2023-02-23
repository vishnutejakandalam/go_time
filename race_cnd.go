/* A Go program to clearly illustrate the Go race conditions. */

package main

// import packages.
import (
	"fmt"
	"time"
)

func main() {
	i := 0
	// Create a simple go routine that prints 10 numbers and increemnts i everytim.
	go func() {
		for i < 10 {
			fmt.Println("This is A", i)
			time.Sleep(500)
			i++
		}
	}()

	// Another routine that has the same i but stops after each iteration
	func() {
		for i < 10 {
			fmt.Println("This is B", i)
			time.Sleep(400)
			i++
		}
	}()

	/* When the second method sleeps for 500 milliseconds, the first go routine starts executing with wrong value of i. Thus giving false output.

	This is called race condition.
	*/

	fmt.Printf("Hello world... \n")
}
