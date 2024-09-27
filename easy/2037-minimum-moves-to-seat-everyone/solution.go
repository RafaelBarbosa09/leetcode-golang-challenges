package main

import (
	"fmt"
	"math"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	total := 0

	for i := 0; i < len(seats); i++ {
		total += int(math.Abs(float64(seats[i] - students[i])))
	}

	return total
}

func main() {
	fmt.Println(minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
}
