package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string = ""
	addr := ""
	fmt.Printf("Enter a name: ")
	fmt.Scanln(&name)
	fmt.Printf("Enter the address: ")
	fmt.Scanln(&addr)
	dict := make(map[string]string)
	dict["name"] = name
	dict["address"] = addr
	json_obj, _ := json.Marshal(dict)
	fmt.Println(dict)
	fmt.Print("The json object is ", string(json_obj))
}
