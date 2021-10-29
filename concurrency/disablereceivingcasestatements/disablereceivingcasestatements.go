/*
In this example, a goroutine is started and using a select tries to receive on
two channels. If a channel is closed it is set to nill. Because nil channels
always block, this has the effect of disabling the associated case statement.
If both channels are set to nil, the goroutine exits because it cannot receive
anything else.
*/
package main

import "fmt"

func main() {
	c1 := make(chan bool)
	c2 := make(chan bool)
	c3 := make(chan bool)

	go func() {
		for {
			select {
			case x, ok := <-c1:
				if !ok {
					c1 = nil
				}
				fmt.Printf("1: %v \n", x)

			case x, ok := <-c2:
				if !ok {
					c2 = nil
				}
				fmt.Printf("2: %v \n", x)
			case x, ok := <-c3:
				if !ok {
					c3 = nil
				}
				fmt.Printf("3: %v \n", x)
			}

			if c1 == nil && c2 == nil && c3 == nil {
				return
			}
		}
	}()

	c1 <- true

	// Disable the c1 case in the above select.
	close(c1)

	c2 <- true
	c2 <- true
	c2 <- true

	c3 <- true
	c3 <- true
}
