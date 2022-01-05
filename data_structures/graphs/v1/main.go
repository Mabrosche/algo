package main

import (
	"fmt"
)

func main() {
	test := &graph{}

	for i := 0; i < 5; i++ {
		test.addVertex(i)
	}

	test.addVertex(0)
	test.addVertex(0)
	test.addVertex(0)
	test.addVertex(0)

	test.print()
}

type graph struct {
	vertices []*vertex
}

type vertex struct {
	val      int
	adjacent []*vertex
}

func (g *graph) addEdges(from, to int) {

}

func (g *graph) addVertex(value int) {
	if contains(g.vertices, value) {
		fmt.Println("Vertex not added because it is an existing key")
	} else {
		vertex := &vertex{
			val: value,
		}

		g.vertices = append(g.vertices, vertex)
	}
}

func (g graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.val)

		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.val)
		}
	}
}

func contains(s []*vertex, value int) bool {
	for _, v := range s {
		if value == v.val {
			return true
		}
	}

	return false
}
