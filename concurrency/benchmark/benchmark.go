package benchmark

import (
	"testing"
)

//BenchmarkBoolInt-4              21384556                49.9 ns/op
//BenchmarkBoolChan-4             23025247                48.2 ns/op
//BenchmarkStructChan-4           22657322                47.0 ns/op
//BenchmarkStructPtrChan-4        20044306                56.4 ns/op

var i int

func BenchmarkBoolInt(b *testing.B) {
	c := make(chan int, b.N)
	for i := 0; i < b.N; i++ {
		c <- i
	}
	close(c)

	for range c {
		i = <-c
	}
}

var bl bool

func BenchmarkBoolChan(b *testing.B) {
	c := make(chan bool, b.N)
	for i := 0; i < b.N; i++ {
		c <- true
	}
	close(c)

	for range c {
		bl = <-c
	}
}

var s struct{}

func BenchmarkStructChan(b *testing.B) {
	c := make(chan struct{}, b.N)
	for i := 0; i < b.N; i++ {
		c <- struct{}{}
	}
	close(c)

	for range c {
		s = <-c
	}
}

var p *struct{}

func BenchmarkStructPtrChan(b *testing.B) {
	c := make(chan *struct{}, b.N)
	for i := 0; i < b.N; i++ {
		c <- nil
	}
	close(c)

	for range c {
		p = <-c
	}
}
