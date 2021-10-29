package main

import "fmt"

// can be used to consume all values from a channel
// https://www.youtube.com/watch?v=SmoM1InWXr0
//https://github.com/nomad-software/go-channel-compendium
func generator(strings chan string) {
	strings <- "1"
	strings <- "2"
	strings <- "3"
	strings <- "4"

	close(strings)
}

func main() {
	strings := make(chan string)

	go generator(strings)

	for s := range strings {
		fmt.Printf("%s", s)
	}

	fmt.Printf("\n")
}
