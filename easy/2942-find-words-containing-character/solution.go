package main

import "fmt"

func getSneakyNumbers(nums []int) []int {
	foundsMap := make(map[int]bool)
	foundsList := []int{}

	for i := 0; i < len(nums); i++ {
		if foundsMap[nums[i]] {
			foundsList = append(foundsList, nums[i])
		}

		foundsMap[nums[i]] = true
	}

	return foundsList
}

func main() {
	fmt.Println(getSneakyNumbers([]int{0, 1, 1, 0}))
}
