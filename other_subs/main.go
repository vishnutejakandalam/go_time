package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Enter the filename:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	names := make([]name, 0, 20)

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		nms := strings.Fields(scanner.Text())
		names = append(names, name{fname: nms[0], lname: nms[1]})
	}

	for _, n := range names {
		fmt.Println("First name:", n.fname)
		fmt.Println("Last name:", n.lname)
	}
	if err != nil {
		log.Fatalln(err)
	}
}
