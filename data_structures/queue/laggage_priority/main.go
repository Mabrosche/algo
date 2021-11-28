package main

type Belt []*Luggage

func main() {
	belt := make(Belt, 0)
	belt.Add(NewLuggage(3, 1, "Elmer Fudd"))
	belt.Add(NewLuggage(3, 1, "Sylvester"))
	belt.Add(NewLuggage(3, 1, "Yosemite Sam"))
	belt.Inspect()

	belt.Add(NewLuggage(3, 2, "Daffy Duck"))
	belt.Inspect()

	belt.Add(NewLuggage(3, 3, "Bugs Bunny"))
	belt.Inspect()

	belt.Add(NewLuggage(100, 2, "Wile E. Coyote"))
	belt.Inspect()
}

type Luggage struct {
	weight    int
	priority  int
	passenger string
}

func NewLuggage(weight int, priority int, passenger string) *Luggage {
	l := Luggage{
		weight:    weight,
		priority:  priority,
		passenger: passenger,
	}
	return &l
}

func (belt *Belt) Add(newLuggage *Luggage) {
	if len(*belt) == 0 {
		*belt = append(*belt, newLuggage)
	} else {
		added := false
		for i, placedLuggage := range *belt {
			if newLuggage.priority > placedLuggage.priority {
				*belt = append((*belt)[:i], append(Belt{newLuggage}, (*belt)[i:]...)...)
				added = true
				break
			}
		}
		if !added {
			*belt = append(*belt, newLuggage)
		}
	}
}
