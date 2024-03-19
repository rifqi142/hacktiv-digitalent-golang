package main

import (
	"hacktiv-digitalent-golang/assignment/assignment-3/handler"
	"net/http"
)

func main() {
	handler.StatusAutoReload()

	http.HandleFunc("/status", handler.StatusHandler)
	http.HandleFunc("/", handler.IndexHandler)

	http.ListenAndServe(":8080", nil)
}