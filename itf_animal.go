package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Snake struct{}
type Bird struct{}
type Cow struct{}

// eat methods for three animals.
func (b Bird) Eat() {
	fmt.Print("worms")
}
func (s Snake) Eat() {
	fmt.Print("mice")
}
func (c Cow) Eat() {
	fmt.Print("grass")
}

// Move methods for three animals
func (b Bird) Move() {
	fmt.Print("fly")
}
func (s Snake) Move() {
	fmt.Print("slither")
}
func (c Cow) Move() {
	fmt.Print("walk")
}

// Speak methods for three animals
func (b Bird) Speak() {
	fmt.Print("peep")
}
func (s Snake) Speak() {
	fmt.Print("hsss")
}
func (c Cow) Speak() {
	fmt.Print("moo")
}

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	animals_list := make(map[string]Animal, 0)
	var data string
	for {
		fmt.Print(">")
		Scanner.Scan()
		data = Scanner.Text()
		set := strings.Split(data, " ")
		if set[0] == "newanimal" {
			name := set[1]
			ani_type := set[2]
			var ani Animal
			if ani_type == "bird" {
				ani = Bird{}
			}
			if ani_type == "snake" {
				ani = Snake{}
			}
			if ani_type == "cow" {
				ani = Cow{}
			}
			animals_list[name] = ani
			fmt.Print("Created it!")
		}
		if set[0] == "query" {
			name := set[1]
			action := set[2]
			obj, ok := animals_list[name]
			if !ok {
				fmt.Print("No animal named ", name)
				continue
			}
			if action == "eat" {
				obj.Eat()
			}
			if action == "move" {
				obj.Move()
			}
			if action == "speak" {
				obj.Speak()
			}

		}
		fmt.Println()
	}

}
