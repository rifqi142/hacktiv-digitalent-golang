package main

// import (
// 	"fmt"
// 	"net/http"
// )

// var PORT = ":8080"

// func main() {
// 	http.HandleFunc("/", greet)

// 	go func() {
// 		if err := http.ListenAndServe(PORT, nil); err != nil {
// 			fmt.Printf("Server error: %v\n", err)
// 		} else {
// 			fmt.Println("Server is running on port", PORT)
// 		}
// 	}()

// 	// Pesan ini akan dicetak setelah server berhasil berjalan
// 	fmt.Println("Server started successfully.")
// }

// func greet (w http.ResponseWriter, r *http.Request) {
// 	msg := "Hello, welcome to Go Learning Center"
// 	fmt.Fprint(w, msg)
// }