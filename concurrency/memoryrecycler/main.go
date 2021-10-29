/*
In this example, a goroutine is started to recycle memory buffers. The give
channel receives old memory buffers and stores them in a list. While the get
channel dispenses these buffers for use. If no buffers are available in the
list, a new one is created.
*/
package main

import "container/list"

func main() {

	give := make(chan []byte)
	get := make(chan []byte)

	go func() {
		var q list.List

		for {
			if q.Len() == 0 {
				q.PushFront(make([]byte, 100))
			}

			e := q.Front()

			select {
			case s := <-give:
				q.PushFront(s)

			case get <- e.Value.([]byte):
				q.Remove(e)
			}
		}
	}()

	// Gets a new buffer from the recycler.
	buffer := <-get

	// Give it back to the recycler.
	give <- buffer

	// Get the recycled buffer again.
	buffer = <-get
}
