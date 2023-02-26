package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal structure
type Animal struct {
	name       string
	food       string
	locomotion string
	sound      string
}

// Eat returns the food the animal eats
func (animal Animal) Eat() string {
	return animal.food
}

// Move returns the way the animal moves
func (animal Animal) Move() string {
	return animal.locomotion
}

// Speak returns the way the animal sounds
func (animal Animal) Speak() string {
	return animal.sound
}

var animals map[string]Animal = map[string]Animal{
	"cow":   Animal{name: "cow", food: "grass", locomotion: "walk", sound: "moo"},
	"bird":  Animal{name: "bird", food: "worms", locomotion: "fly", sound: "peep"},
	"snake": Animal{name: "snake", food: "mice", locomotion: "slither", sound: "hsss"},
}

var actions map[string]func(Animal) string = map[string]func(Animal) string{
	"eat":   func(animal Animal) string { return animal.Eat() },
	"move":  func(animal Animal) string { return animal.Move() },
	"speak": func(animal Animal) string { return animal.Speak() },
}

func processCommand(animalType, actionType string) (string, bool) {
	animal, animalPresent := animals[animalType]
	action, actionPresent := actions[actionType]

	if animalPresent && actionPresent {
		return action(animal), true
	}

	return "", false
}

func printUsage() {
	fmt.Println("usage: [bird|cow|snake eat|move|speak]|[bye]")
	fmt.Println("example: > bird eat")
	fmt.Println("example: > bye")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			words := scanner.Text()
			if words == "bye" {
				break
			}

			commands := strings.Split(words, " ")

			if len(commands) != 2 {
				printUsage()
				continue
			}

			animalType := commands[0]
			actionType := commands[1]
			result, ok := processCommand(animalType, actionType)

			if ok {
				fmt.Println(result)
			} else {
				printUsage()
			}
		}
	}

	fmt.Println("bye!")
}
