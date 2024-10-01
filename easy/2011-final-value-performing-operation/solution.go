package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	acc := 0

	for i := 0; i < len(operations); i++ {
		if operations[i][1] == '-' {
			acc--
			continue
		}

		acc++
	}

	return acc
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}
