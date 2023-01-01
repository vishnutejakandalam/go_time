package main

import "fmt"

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + s0
	}
}

func main() {
	var acc, ini_vlc, ini_dspl float64
	fmt.Println("Enter the acceleration: ")
	fmt.Scanf("%f", &acc)
	fmt.Println("Enter the initial velocity: ")
	fmt.Scanf("%f", &ini_vlc)
	fmt.Println("Enter the initial displacement: ")
	fmt.Scanf("%f", &ini_dspl)
	function := GenDisplaceFn(acc, ini_vlc, ini_dspl)
	var time float64
	fmt.Println("Enter the time: ")
	fmt.Scanf("%f", &time)
	fmt.Println(function(time))
}
