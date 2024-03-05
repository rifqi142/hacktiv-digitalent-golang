package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/person", personHandler)
	
	http.ListenAndServe("localhost:8082", nil)
}

var persons = make([]Person, 0)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}


func addPerson (w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)

	if err != nil {
		response := Response {
			Success: false,
			Error: err.Error(),
		}
		jsonResponse(response, http.StatusInternalServerError, w)
		return
	}

	var newPerson Person
	
	err = json.Unmarshal(data, &newPerson)
	if err != nil {
		response := Response {
			Success: false,
			Error: err.Error(),
		}
		jsonResponse(response, http.StatusBadRequest, w)
		return
	}

	persons = append(persons, newPerson)

	var response Response = Response {
		Success: true,
		Data: newPerson,
	}

	jsonResponse(response, http.StatusOK, w)

}

func getAllPerson(w http.ResponseWriter, r *http.Request) {
	var response Response = Response {
		Success: true,
		Data: persons,
	}
	jsonResponse(response, http.StatusOK, w)
}



func personHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllPerson(w, r)
	case http.MethodPost:
		addPerson(w, r)

	default:
		response := Response{
			Success: false,
			Error:   "not found",
		}

		jsonResponse(response, http.StatusNotFound, w)
	}
}

func jsonResponse(r Response, httpCode int, w http.ResponseWriter) {
	result, err := json.Marshal(r)
	if err != nil {
		fmt.Println("error marshalling", err)
		http.Error(w, fmt.Sprintf("error json encoding %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(result)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ini dari fmt.Println") // io server
	w.Write([]byte("hello world"))
}
