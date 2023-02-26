package main

import "fmt"

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	fmt.Println("Enter acceleration:")
	var a float64
	fmt.Scan(&a)

	fmt.Println("Enter initial velocity:")
	var vo float64
	fmt.Scan(&vo)

	fmt.Println("Enter initial displacement:")
	var so float64
	fmt.Scan(&so)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Println("Enter time:")
	var t float64
	fmt.Scan(&t)

	displacement := fn(t)
	fmt.Println("Displacement after", t, "seconds:", displacement)
}