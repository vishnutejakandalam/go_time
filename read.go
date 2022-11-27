package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	list := make([]name, 0)
	file_name := ""
	fmt.Println("Enter the text file to read:  ")
	fmt.Scan(&file_name)
	f, _ := os.Open(file_name)
	s := bufio.NewScanner(f)
	for s.Scan() {
		set := strings.Split(s.Text(), " ")
		var fullname name
		fullname.fname = set[0]
		fullname.lname = set[1]
		list = append(list, fullname)
		// list = append(list)
	}
	for i, l := range list {
		fmt.Println(i, l)
	}
}
