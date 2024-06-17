package main

import "fmt"

func main() {
	l1, l2, l3 := linked()
	fmt.Println("Linked     :", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("No Link     :", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap Link     :", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link    :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link    :", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link :", a1, a2)
}

func linked() (int, int, int) {
	// slice of int, initialized with 5 values
	s1 := []int{1, 2, 3, 4, 5}
	// simple copy variable of s1
	s2 := s1
	// create a new slice with data from the first slice with slice range operation
	s3 := s1[:]

	// change data in s1 slice to s2 slice
	s1[3] = 99
	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	// creating a slice this way, will have its length and capacity autoset to the number of values passed in.
	// in this case,
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	// appending this will be pass s1's capacity, so the underlying array will be expanded. When expanding, the values
	// from the underlying array into a new array so s2 will still point to the initial underlying array, while s1 will point to the
	// new expanded array
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	// creating a new slice with length of 5 and underlying array capacity of 10
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	// appending a value to s1. this will change its length to 6, but still under the capacity of the underlying array
	// so no expanding is needed
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5

	s2 := s1
	// going beyong capacity of 10
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	// copy the elements of s1 into s2. copied is the value of the number of elements copied over
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	// setting a new hidden array when we add {} as part of the parameters being passed in
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}
