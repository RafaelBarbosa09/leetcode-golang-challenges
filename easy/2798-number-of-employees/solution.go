package main

import "fmt"

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	count := 0
	for _, item := range hours {
		if item >= target {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 2))
}
