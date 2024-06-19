package main

import "fmt"

type name string // custom type "name" of type string

type location struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

// embedding name, location, and size into the dot struct
type dot struct {
	name
	location
	size
}

func main() {
	dots := getDots()
	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v: %#v\n", i+1, dots[i])
	}
}

func getDots() []dot {
	// dots with zeroed out values
	var dot1 dot
	/*
		When embedding, the fields from the sub structs (location, size), gets promoted to the root level of dot struct.
		Since "name" is just a string, no promotion is needed. If the dot struct had a field on its root with the same
		name as a sub structs field, then to access sub struct's field, you will need to traverse the type path
		(dotVal.[type].[field] ex. dot2.location.commonField).
	*/
	dot2 := dot{}
	dot2.name = "A"
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	// When initializing a struct, promotion cannot be used.
	dot3 := dot{
		name: "B",
		location: location{
			x: 13,
			y: 27,
		},
		size: size{
			width:  5,
			height: 7,
		},
	}

	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 209
	dot4.size.width = 87
	dot4.size.height = 43

	return []dot{dot1, dot2, dot3, dot4}
}

// Note: embedding is not all that common as most Go devs prefer just to have structs as a field to the parent struct
