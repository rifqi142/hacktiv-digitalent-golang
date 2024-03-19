package main

import (
	"fmt"

	"encoding/json"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `[
	{
		"full_name": "Muhammad Rifqi",
		"email": "rifqi@gmail.com",
		"age": 23
	},
	{
		"full_name": "Fadlan Amrullah",
		"email": "fadlan@gmail.com",
		"age": 23
	}
	]
	`

	// !!decoding biasa
	// var result Employee

	// !! decoding json to map
	// var result map[string]interface{}

	// !! decoding json to empty interface
	// var result interface{}

	// decoding JSON array to slice of struct
	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println("Full Name:", result.FullName)
	// fmt.Println("Email:", result.Email)
	// fmt.Println("Age:", result.Age)

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}
