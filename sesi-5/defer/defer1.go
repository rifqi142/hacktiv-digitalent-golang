package main

import "fmt"

func main() {
	//  akan dijakanlan apabila main function selesai dijalankan
	// defer fmt.Println("defer function starts to execute")
	// fmt.Println("Hello everyone~")
	// fmt.Println("Welcome back to go Learning Center!!")

	var a int = 5
	defer func () {
		fmt.Println("ini dari defer", a)
	} ()

	a = 456
	fmt.Println("ini dari main", a)

	for i := 0; i < 10; i++ {
		defer func () {	
			fmt.Println(i)
		} ()
	}
} 