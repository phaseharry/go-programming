package main

import "fmt"

/*
	type Number interface {
		int | float64
	}

is a type constraints we created for our generic function.
"Type Constraints" in Go generics refer to interfaces that define sets of types.
Type constraints specifies the requirements or capabilities that a type parameter must satisfy
when working with generic functions or types. By declaring type constraints, this specifies to
the Go compiler the expected behavior of the type parameter and type checks can be performed
during compile time.
*/
type Number interface {
	int | float64
}

// parameterizing this generic function with the declared type constraint above
func findMaxGeneric[Num Number](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		currentNum := nums[i]
		if currentNum > max {
			max = currentNum
		}
	}
	return max
}

func main() {
	maxGenericInt := findMaxGeneric([]int{1, 32, 5, 8, 10, 11})
	fmt.Printf("max generic int %v\n", maxGenericInt)

	maxGenericFoat := findMaxGeneric([]float64{1.5, 32.32, 5.1, 8.8, 10.213, 11.561})
	fmt.Printf("max generic float %v\n", maxGenericFoat)
}
