package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	bigger := 0
	for i := 0; i < len(accounts); i++ {
		acc := 0
		for j := 0; j < len(accounts[i]); j++ {
			acc += accounts[i][j]
		}

		if acc > bigger {
			bigger = acc
		}
	}

	return bigger
}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
}
