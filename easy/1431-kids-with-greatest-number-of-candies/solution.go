package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}

	list := []bool{}
	for j := 0; j < len(candies); j++ {
		list = append(list, candies[j] + extraCandies >= maxCandies)
	}

	return list
}

func main() {
	fmt.Println(kidsWithCandies([]int{2,3,5,1,3}, 3))
}
