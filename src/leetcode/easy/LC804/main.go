package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}

var MorseArray = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	m := map[string]int{}
	for _, w := range words {
		var sb strings.Builder
		for _, c := range w {
			sb.WriteString(MorseArray[c-'a'])
		}
		m[sb.String()] = 0
	}
	return len(m)
}
