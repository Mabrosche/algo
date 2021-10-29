package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 15
	c <- 34
	c <- 65
	close(c)

	// drains the buffered data
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)

	// starts returning the ZERO value
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
}
