package main

import "fmt"

type Int struct {
	Slice []int
	Len   int
}

func (q *Int) Enqueue(i int) {
	q.Slice = append(q.Slice, i)
	q.Len++
}

func (q *Int) Dequeue() int {
	res := q.Slice[0]
	q.Slice = q.Slice[1:len(q.Slice)]
	q.Len--
	return res
}

func (q Int) String() string {
	return fmt.Sprint(q.Slice)
}

func (q Int) IsEmpty() bool {
	return q.Len == 0
}
