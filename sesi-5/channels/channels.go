package main

import "fmt"

/*
 c chan <- string = c hanya dapat digunakan untuk mengirim data
 c <-chan string = c hanya dapat digunakan untuk menerima data
*/

// func main() {
// 	c := make(chan string)

// 	students := []string{"Airell", "Rifqi", "Ali"}

// 	for _, v := range students {
// 		go introduce(v,c)
// 	}

// 	for i := 1; i<=3; i++ {
// 		print(c)
// 	}

// 	close (c)
// }

// func print(c <- chan string) {
// 	fmt.Println(<-c)
// }

// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("Hello, my name is %s", student)

// 	c <- result
// }

// func main () {
// 	c := make(chan int)

// 	go channel(c)
// 	c <- 10
// 	fmt.Println(<-c) // 20

// 	go parameter(c, 11)
// 	fmt.Println(<-c) // 21
// }

// func channel(c chan int) {
// 	input := <-c  // input = 10
// 	c <- 10 + input // c <- 20
// }

// func parameter(c chan int, n int) {
// 	input := n // 11
// 	c <- 10 + input // c <- 21
// }

func main () {
	c1 := make(chan int)
	c2 := make(chan int)

	go f1(c1)
	go f2(c2)

	for i := 0; i < 2; i++   {
		select {
		case <-c1 :
			fmt.Println("Channel 1 Sampai")
		case <-c2 :
			fmt.Println("Channel 2 Sampai")
		}
	}
}


func f1 (c chan int) {
	c <- 333
}

func f2 (c chan int) {
	c <- 222
}
