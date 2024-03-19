package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// create body request
	data := map[string]string{
		"title":  "rifqi",
		"body":   "setiawan",
		"userId": "1",
	}
	// Convert map to json
	requestJson, err := json.Marshal(data)

	client := &http.Client{}
	if err != nil {
		log.Fatal(err)
	}
	// Create new request
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")
	
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}