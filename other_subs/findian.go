package main

import (
    "fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	fmt.Println( "Please enter a string:" )
	
	var s string

	scanner := bufio.NewScanner( os.Stdin )
	scanner.Scan()
	s += scanner.Text()

	s = strings.ToLower( s )
	last := s[ len( s ) - 1 ]

	if strings.Contains( s, "a" ) && s[0] == 'i' && last == 'n' {
		fmt.Println("Found!")

	} else {
		fmt.Println("Not Found!")

	}

}