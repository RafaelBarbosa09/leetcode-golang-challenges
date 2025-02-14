package main

import "fmt"

func main() {
	fmt.Println(subsetXORSum([]int{1, 3}))
}

func backtrack(nums []int, index int, currentXOR int) int {
	if index == len(nums) {
		return currentXOR
	}

	withoutCurrent := backtrack(nums, index+1, currentXOR)

	withCurrent := backtrack(nums, index+1, currentXOR^nums[index])

	return withoutCurrent + withCurrent
}

func subsetXORSum(nums []int) int {
	return backtrack(nums, 0, 0)
}
