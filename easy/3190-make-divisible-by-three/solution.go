package main

import "fmt"

func minimumOperations(nums []int) int {
	operations := 0

	for _, item := range nums {
		if item % 3 != 0 {
			operations++
		}
	}

	return operations
}

func main() {
	fmt.Println(minimumOperations([]int{1, 2, 3, 4}))
}
