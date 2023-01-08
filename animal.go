package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	if a.name == "cow" {
		fmt.Println("grass")
	} else if a.name == "bird" {
		fmt.Println("worms")
	} else if a.name == "snake" {
		fmt.Println("mice")
	}
}

func (a *Animal) Move() {
	if a.name == "cow" {
		fmt.Println("walk")
	} else if a.name == "bird" {
		fmt.Println("fly")
	} else if a.name == "snake" {
		fmt.Println("slither")
	}
}

func (a *Animal) Speak() {
	if a.name == "cow" {
		fmt.Println("moo")
	} else if a.name == "bird" {
		fmt.Println("peep")
	} else if a.name == "snake" {
		fmt.Println("hsss")
	}
}

func main() {
	var data string
	Scanner := bufio.NewScanner(os.Stdin)
	a := Animal{}
	for {
		fmt.Printf(">")
		Scanner.Scan()
		data = Scanner.Text()
		set := strings.Split(data, " ")
		animal := strings.ToLower(set[0])
		action := strings.ToLower(set[1])
		a.name = animal
		if action == "eat" {
			a.Eat()
		} else if action == "move" {
			a.Move()
		} else if action == "speak" {
			a.Speak()
		} else {
			fmt.Println("WRONG CHOICE..")
			os.Exit(0)
		}
	}
}
