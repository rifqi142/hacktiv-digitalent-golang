package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "John Doe", Age: 25, Division: "IT"},
	{ID: 2, Name: "Rick Ruer", Age: 30, Division: "HR"},
	{ID: 3, Name: "Budiman", Age: 35, Division: "Finance"},
	{ID: 4, Name: "Juliano", Age: 40, Division: "Marketing"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application is listening on port", PORT)

	http.ListenAndServe(PORT, nil)
}



func getEmployees(w http.ResponseWriter, r *http.Request) {
	// without html
	// w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// without html
		// json.NewEncoder(w).Encode(employees)
		// return

		// with html
		tpl, err := template.ParseFiles("./template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID: len(employees) + 1,
			Name: name,
			Age: convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Method not allowed", http.StatusBadRequest)
}