package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}) == "world")
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}) == "apple")
}

type Dic struct {
	m map[byte]*Dic
}

func longestWord1(words []string) string {
	sort.Strings(words)
	length := 0
	var str string
	m := Dic{map[byte]*Dic{}}
	for _, s := range words {
		l := len(s)
		if l == 1 {
			m.m[s[0]] = &Dic{map[byte]*Dic{}}
			if length < l {
				length = 1
				str = s
			}
			continue
		}
		v := m.m[s[0]]
		if v == nil {
			continue
		}
		for i := 1; i < l; i++ {
			if i == l-1 {
				v.m[s[i]] = &Dic{map[byte]*Dic{}}
				if length < l {
					length = l
					str = s
				}
			} else {
				v = v.m[s[i]]
				if v == nil {
					break
				}
			}
		}
	}
	return str
}

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			// Make sure words with smaller lexicographical order come first
			// when we iterate over the words in reverse order
			return words[i] > words[j]
		}
		return len(words[i]) < len(words[j])
	})

	for i := len(words) - 1; i >= 0; i-- {
		w := words[i]
		allFound := true
		for l := 1; l < len(w); l++ {
			target := w[:l]
			found := sort.Search(i, func(j int) bool {
				if len(target) == len(words[j]) {
					return words[j] <= target
				}
				return len(words[j]) > len(target)
			})
			allFound = found < i && words[found] == target
			if !allFound {
				break
			}
		}
		if allFound {
			return w
		}
	}
	return ""
}
