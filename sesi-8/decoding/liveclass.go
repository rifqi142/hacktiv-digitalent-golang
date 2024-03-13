package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	var a = `
	{
		"name": "John Doe",
		"age": 30
	}
	`

	var p Person
	var m map[string]any
	var x any

	json.Unmarshal([]byte(a), &p)
	fmt.Println(p)

	json.Unmarshal([]byte(a), &m)
	fmt.Println(m, m["age"])

	json.Unmarshal([]byte(a), &x)
	fmt.Println(x.(map[string]any)["name"])

	fmt.Println(json.Marshal(m))
}