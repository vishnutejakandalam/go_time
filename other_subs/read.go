package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	maxLength = 20
)

type Name struct {
	fName string
	lName string
}

func (n *Name) Set(firstName string, lastName string) {
	n.fName = firstName
	n.lName = lastName
	if len(firstName) > maxLength {
		n.fName = firstName[:maxLength]
	}
	if len(lastName ) > maxLength {
		n.lName = lastName[:maxLength]
	}
}

func (n *Name) Get() string {
	return n.fName + " " + n.lName
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	var fileName string
	fmt.Println("Please enter the name of the file you wish to read:")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	var namesSlice []Name

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var n Name;
		line := scanner.Text()
		ns := strings.Fields(line)
		fName := ns[0]
		lName := ns[1]
		n.Set(fName, lName)
		namesSlice = append(namesSlice, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, n := range namesSlice {
		fmt.Println(n.Get())
	}
}