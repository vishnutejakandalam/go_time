package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter address: ")
	fmt.Scan(&address)

	//map
	map_1 := map[string]string{
		"name":    name,
		"address": address,
	}
	//json
	jsonString, err := json.Marshal(map_1)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonString))
}
