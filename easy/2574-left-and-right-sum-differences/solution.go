package main

import (
	"fmt"
	"math"
)

func leftRightDifference(nums []int) []int {
	n := len(nums)
	leftSum := make([]int, n)
	rightSum := make([]int, n)

	for i := 0; i < n; i++ {
		if i >= 1 {
			leftSum[i] = leftSum[i-1] + nums[i-1]
			rightSum[n-i-1] = nums[n-i] + rightSum[n-i]
		}
	}

	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = int(math.Abs(float64(leftSum[i] - rightSum[i])))
	}

	return list
}

func leftRightDifferenceSecondApproach(nums []int) []int {
	list := []int{}

	var left, right int
	for _, x := range nums {
		right += x
	}

	for _, x := range nums {
		right -= x
		list = append(list, int(math.Abs(float64(left-right))))
		left += x
	}

	return list
}

func main() {
	fmt.Println(leftRightDifference([]int{10, 4, 8, 3}))
	fmt.Println(leftRightDifferenceSecondApproach([]int{10, 4, 8, 3}))
}
