package main

import "fmt"

func shuffle(nums []int, n int) []int {
	var list []int

	for i := 0; i < n; i++ {
		list = append(list, nums[i])
		list = append(list, nums[i+n])
	}

	return list
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
