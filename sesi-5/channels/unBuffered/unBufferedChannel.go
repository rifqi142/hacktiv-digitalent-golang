package main

import (
	"fmt"
	"time"
)

func main () {
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channels")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	} (c1)

	fmt.Println("main goroutine sleeps for 2 second")
	time.Sleep(time.Second * 2)

	fmt.Println("main gorooutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	close (c1)
	time.Sleep(time.Second)

}