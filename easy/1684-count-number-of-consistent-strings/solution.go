package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	allowedMap := make(map[rune]bool)
	for _, char := range allowed {
		allowedMap[char] = true
	}

	consistentCount := 0

	for _, word := range words {
		isConsistent := true

		for _, char := range word {
			if !allowedMap[char] {
				isConsistent = false
				continue
			}
		}

		if isConsistent {
			consistentCount++
		}
	}

	return consistentCount
}

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
}
