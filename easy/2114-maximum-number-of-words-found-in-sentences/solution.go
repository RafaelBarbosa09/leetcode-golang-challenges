package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))
}

func mostWordsFound(sentences []string) int {
	greater := 0
	for _, sentence := range sentences {
		replace := strings.ReplaceAll(sentence, " ", ",")
		split := strings.Split(replace, ",")

		if len(split) > greater {
			greater = len(split)
		}
	}

	return greater
}
