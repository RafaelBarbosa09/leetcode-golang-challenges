package main

import "fmt"

func getConcatenation(nums []int) []int {
	var list []int

	for i := 0; i < 2; i++ {
		list = append(list, nums...)
	}

	return list
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
}
