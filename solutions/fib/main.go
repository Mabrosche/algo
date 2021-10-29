package main

import (
	"fmt"
)

func main() {
	res := fib(91)

	fmt.Println(res)
}

// 		n: 1, 2, 3, 4, 5, 6, 7,  8,  9
// fib(n): 1, 1, 2, 3, 5, 8, 13, 21, 34
func fibReq(n int) int {
	if n <= 2 {
		return 1
	}

	return fibReq(n-1) + fibReq(n-2)
}

//
//func fibDP(n int) int {
//	// Memoization
//	mem := make([]int, n+1)
//	if mem[n]
//
//	return mem[n]
//}

var m = make(map[int]int)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if _, ok := m[n]; ok {
		return m[n]
	}

	if n <= 2 {
		return 1
	}

	m[n] = fib(n-1) + fib(n-2)

	return m[n]
}
