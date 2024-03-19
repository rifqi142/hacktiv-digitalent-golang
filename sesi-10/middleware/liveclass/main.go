package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", logMiddleware(authMiddleware(helloHandler)))

	http.ListenAndServe("localhost:8082", nil)
}

func logMiddleware (next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		 fmt.Println("Log: ", r.URL.Path)
		next(w, r) 
	}
}

func authMiddleware (next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		 
		 username, password, ok := r.BasicAuth()
		 if !ok {
			w.Write([]byte("Not Authorized"))
			return
		 }

		 if username != "admin" || password != "admin" {
			w.Write([]byte("username or password is wrong"))
			return
		 }
		next(w, r) 
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}