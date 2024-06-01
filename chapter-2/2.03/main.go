package main

import "fmt"

func main() {
	// bubble sort
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	numsSorted := 0
	sorted := false
	fmt.Printf("Before sort %#v\n", nums)
	for {
		sorted = true
		unsortedCount := len(nums) - numsSorted - 1
		for i := 0; i < unsortedCount; i++ {
			currentNum := nums[i]
			nextNum := nums[i+1]
			// bubbling up currentNum if it's greater than nextNum
			if currentNum > nextNum {
				nums[i], nums[i+1] = nextNum, currentNum
				sorted = false
			}
		}
		numsSorted += 1
		if sorted {
			break
		}
	}
	fmt.Printf("After sort %#v\n", nums)

}
