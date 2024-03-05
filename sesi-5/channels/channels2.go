package main

// import "fmt"

/*
 c chan <- string = c hanya dapat digunakan untuk mengirim data
 c <-chan string = c hanya dapat digunakan untuk menerima data
*/

// func main() {
// 	c := make(chan string)

// 	students := []string{"Airell", "Rifqi", "Ali"}

// 	for _, v := range students {
// 		go func (student string) {
// 			fmt.Println("Student", student)
// 			result := fmt.Sprintf("Hello, my name is %s", student)
// 			c <- result
// 		} (v)
// 	}

// 	for i := 1; i <= 3; i++ {
// 		print(c)
// 	}

// 	close (c)
//  }

//  func print (c chan string) {
// 	fmt.Println(<-c)
//  }

