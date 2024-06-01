package main

import (
	"fmt"
	"time"
)

func main() {

	var count1 *int    // empty pointer that doesn't to point anywhere (nil)
	count2 := new(int) // allocates memory and returns an address that's allowed to store a value of type int

	countTemp := 5
	count3 := &countTemp // stores pointer that points to memory location of countTemp variable

	t := &time.Time{}

	// need to check if a pointer actually points anywhere before we dereference for the value otherwise, it will break
	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("t: %#v\n", *t)
		fmt.Printf("t: %#v\n", t.String())
	}
}
