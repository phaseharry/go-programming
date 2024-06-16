package main

import "fmt"

func main() {
	s1, s2, s3 := generateSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
}

func generateSlices() ([]int, []int, []int) {
	// declaring an empty slice s1
	var s1 []int
	// declaring a slice s2 where it has a length of 10, meaning it has 10 slots of memory filled out with the zero value
	// for int type (0 in this case)
	s2 := make([]int, 10)
	// same as above except the underlying array for this slice has 50 slots of contigous memory to be filled before it needs
	// to be expanded when current slice hits its capacity
	s3 := make([]int, 10, 50)

	return s1, s2, s3
}
