package main

import "fmt"

func main() {
	fmt.Println(stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
	fmt.Println(stringMatching([]string{"mass", "as", "hero", "superhero"}))
	fmt.Println(stringMatching([]string{"leetcode", "et", "code"}))
	fmt.Println(stringMatching([]string{"blue", "green", "bu"}))
}
func stringMatching(words []string) []string {
	var result []string
	m := map[string]struct{}{}
	for index, w := range words {
		for index2, w2 := range words {
			if index == index2 || len(w) <= len(w2) {
				continue
			}
			length1 := len(w)
			length2 := len(w2)
			for i := 0; i <= length1-length2; i++ {
				j := 0
				for ; j < length2; j++ {
					if w[i+j] != w2[j] {
						break
					}
				}
				if j == length2 {
					m[w2] = struct{}{}
				}
			}
		}
	}
	for k := range m {
		result = append(result, k)
	}
	return result
}
