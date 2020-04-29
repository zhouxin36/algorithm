package main

import "fmt"

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach") == 6)
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr") == 10)
}
func countCharacters(words []string, chars string) int {
	chm := map[byte]int{}
	wom := map[byte]int{}
	for _, c := range chars {
		chm[byte(c)]++
	}
	var count int
	for _, word := range words {
		var flag bool
		for _, w := range word {
			v := byte(w)
			wom[v]++
			if wom[v] > chm[v] {
				flag = true
				break
			}
		}
		if !flag {
			count += len(word)
		}
		wom = map[byte]int{}
	}
	return count
}
