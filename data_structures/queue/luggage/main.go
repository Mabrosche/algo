package main

import "fmt"

func main() {
	belt := &Belt{}
	belt.Add(NewLuggage(3, "Elmer Fudd"))
	belt.Add(NewLuggage(5, "Sylvester"))
	belt.Add(NewLuggage(2, "Yosemite Sam"))
	belt.Add(NewLuggage(10, "Daffy Duck"))
	belt.Add(NewLuggage(1, "Bugs Bunny"))

	fmt.Println("Belt:", belt, "Length:", len(*belt))
	first := belt.Take()
	fmt.Println("First luggage:", first)
	fmt.Println("Belt:", belt, "Length:", len(*belt))
}

type Belt []*Luggage

type Luggage struct {
	weight    int
	passenger string
}

func NewLuggage(weight int, passenger string) *Luggage {
	return &Luggage{
		weight:    weight,
		passenger: passenger, // just as an identifier
	}
}

func (belt *Belt) Add(newLuggage *Luggage) {
	*belt = append(*belt, newLuggage)
}

func (belt *Belt) Take() *Luggage {
	first, rest := (*belt)[0], (*belt)[1:]
	*belt = rest
	return first
}
