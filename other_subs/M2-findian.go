package main

import (
	"fmt"
	"strings"
)

func main() {
	var userData = ""

	fmt.Println(
		"Hi! Please enter a string and I will tell you if it has the letters " +
			"I am looking for.")
	fmt.Scan(&userData)
	userDataLower := strings.ToLower(userData)

	startsWithI := strings.HasPrefix(userDataLower, "i")
	containsAnA := strings.Contains(userDataLower, "a")
	endsWithN := strings.HasSuffix(userDataLower, "n")

	if startsWithI && containsAnA && endsWithN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
