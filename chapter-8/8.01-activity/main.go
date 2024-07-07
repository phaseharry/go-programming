package main

import (
	"fmt"
)

func main() {
	intVals := []int{1, 32, 5, 8, 10, 11}
	floatVals := []float64{1.2, 32.5, 5.8, 8.1, -10.41, 11.9}
	/*
		When passing in intVals and floatVals, we are passing our "type arguments" that will provide the type
		that getMinGeneric will be instantiated with and ran with.
	*/
	minInt := getMinGeneric(intVals)
	minFloat := getMinGeneric(floatVals)
	fmt.Printf("minInt: %v, minFloat: %v\n", minInt, minFloat)
}

/*
Generics allow for functions to take in multiple types and functions to be used
for difference types so we don't need to create a new function to support int and float64 seperately.
To do so, we need to type parameterize our generic functions so Go's compiler will understand that it
can take in multiple types.

In this case, we're type parameterizing getMinGenerics to have a "Num" generic type that can take in
either int or float64 values.

Note: Type parameters are usually uppercase to emphasize that they are indeed types.
*/
func getMinGeneric[Num int | float64](nums []Num) Num {
	var minNum Num = nums[0]
	for _, val := range nums {
		if val < minNum {
			minNum = val
		}
	}
	return minNum
}
