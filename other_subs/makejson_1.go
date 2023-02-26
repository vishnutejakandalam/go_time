package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	fmt.Printf("enter name\n")
	fmt.Scan(&name)

	var address string
	fmt.Printf("enter address\n")
	fmt.Scan(&address)

	var m map[string]string
	m = make(map[string]string)
	m["name"] = name
	m["address"] = address

	jsonString, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonString))

}
