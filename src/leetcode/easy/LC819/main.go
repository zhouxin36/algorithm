package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}
func mostCommonWord(paragraph string, banned []string) string {
	m := map[string]struct{}{}
	for _, b := range banned {
		m[b] = struct{}{}
	}
	var letters []string
	var letter strings.Builder
	for _, p := range paragraph {
		if p >= 'A' && p <= 'Z' {
			p += 32
			letter.WriteByte(byte(p))
		} else if p >= 'a' && p <= 'z' {
			letter.WriteByte(byte(p))
		} else {
			if letter.Len() > 0 {
				letters = append(letters, letter.String())
				letter.Reset()
			}
		}
	}
	if letter.Len() > 0 {
		letters = append(letters, letter.String())
		letter.Reset()
	}
	m2 := map[string]int{}
	count := 0
	var result string
	for _, l := range letters {
		if _, ok := m[l]; !ok {
			m2[l]++
			if count < m2[l] {
				count = m2[l]
				result = l
			}
		}
	}
	return result
}
