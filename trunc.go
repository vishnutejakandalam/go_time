package main

import "fmt" 

func main(){
var val float32

fmt.Printf("Enter a floating point number.")
fmt.Scanf("%f", &val)

fmt.Printf("The integer equivalant of %f is %d \n", val, int(val))

}