package main

import "fmt"

func main() {
    var number float32

    fmt.Println("Please enter a float number: ")

    _, err := fmt.Scan(&number)

    if err != nil {
        panic("Error reading user input")
    }

    fmt.Printf("The truncated number is: %v", int32(number))
}
