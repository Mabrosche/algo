package concurrency

import (
	"time"
)

//https://www.youtube.com/watch?v=SmoM1InWXr0
//https://github.com/nomad-software/go-channel-compendium
//https://github.com/cloudflare/jgc-talks/tree/master/GopherCon
//http://www.slideshare.net/cloudflare/a-channel-compendium
//https://gist.github.com/santacruz123/90fe030829a8f8c3b7fb429b429da4f8

//https://itnext.io/explain-to-me-go-concurrency-worker-pool-pattern-like-im-five-e5f1be71e2b0
//https://itnext.io/refactor-go-worker-pool-a-way-around-to-the-sync-package-7d45b1afb768

func TryReceive(c <-chan int) (data int, more, ok bool) {

	select {
	case data, more = <-c:
		return data, more, ok
	default: // process when c is blocking
		return 0, true, false
	}
}

func TryReceiveWithTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, ok
	case <-time.After(duration): // time.After() returns a channel
		return 0, true, false
	}
}

func FanOut(In <-chan int, OutA, OutB chan<- int) {
	for data := range In { // receive until closed
		select { // send to first non-blocking channel
		case OutA <- data:
		case OutB <- data:
		}
	}
}

func IntoCh(conn chan int, data int) {
	conn <- data
}

func OutFromCh(conn chan int) int {
	return <-conn
}
