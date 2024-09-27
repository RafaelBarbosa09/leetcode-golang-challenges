package main

import "fmt"

func nountPairs(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i < j && nums[i]+nums[j] < target {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(nountPairs([]int{-1, 1, 2, 3, 1}, 2))
}
