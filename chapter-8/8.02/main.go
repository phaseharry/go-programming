package main

import "fmt"

/*
More complex example using an Go standard interface "comparable". Any type that can be compared with operators like
"==", "!=", etc. implements the "comparable interface". In our example, we used strings to be the key in our map
that type parameter "K" represents. Strings implements the "comparable interface" so it works for this generic function.
K would also support int / float values since they also implement the "comparable" interface.
*/
func FindLargestRanchStock[K comparable, V int | float64](m map[K]V) K {
	var stock V = 0 // will get type inferenced to int or float64 when a value is passed in
	var name K
	for k, v := range m {
		if v > stock {
			stock = v
			name = k
		}
	}
	return name
}

/*
Go compiler infers the type we want to use from the functions arguments as they're passed in (Type inference).
The compiler will deduce type arguments from type parameter. Type inference can succeed or fail. If it fails,
the compiler will thrown an error and tell you where it failed. If type inference fails, you can declare the
type being passed in when you call the function
*/
func main() {
	animalStock := map[string]int{
		"chicken": 5,
		"cattle":  20,
		"horses":  4,
	}
	miscStock := map[string]float64{
		"hay":        5.5,
		"feed":       1.2,
		"fertilizer": 4.5,
	}
	// go will infer the type in this case
	largestStockOnRanchInt := FindLargestRanchStock(animalStock)
	fmt.Printf("The largest stocked int item on the ranch is %s\n", largestStockOnRanchInt)

	/*
		Can also declare the types as it's being called like so:
		this will explicity tell go to parameterize the FindLargestRanchStock generic function with string & int value.
		Sometimes you need to do this because due to Go being unable to infer the types being passed in.
		Some reasons include ambigous types, multiple type parameters, and chained calls
		- chained calls: scenario where types need to flow to multiple functions.
	*/
	largestStockOnRanchInt = FindLargestRanchStock[string, int](animalStock)

	largestStockOnRanchFloat := FindLargestRanchStock(miscStock)
	fmt.Printf("The largest stocked float64 item on the ranch is %s\n", largestStockOnRanchFloat)
}
