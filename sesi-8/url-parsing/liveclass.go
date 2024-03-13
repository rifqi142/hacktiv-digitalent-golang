package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "https://www.google.com/search?q=golang"
	u, err := url.Parse(urlString)
	fmt.Println(u, err)

	fmt.Println("ini host: ", u.Host)
	fmt.Println("ini scheme: ", u.Scheme)
	fmt.Println("ini path: ", u.Path)
	fmt.Println("ini query: ", u.Query())
}
