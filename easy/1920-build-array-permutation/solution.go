package main

import "fmt"

func buildArray(nums []int) []int {
	var list []int
	for i := 0; i < len(nums); i++ {
		list = append(list, nums[nums[i]])
	}

	return list
}

func main() {
	fmt.Println(buildArray([]int{5, 0, 1, 2, 3, 4}))
}
