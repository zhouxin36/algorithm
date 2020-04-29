package main

import "fmt"

func main() {
	fmt.Println(keyboard("abcdefghijklmnopqrstuvwxyz", "cba") == 4)
	fmt.Println(keyboard("pqrstuvwxyzabcdefghijklmno", "leetcode") == 73)
}
func keyboard(keyboard, word string) int {
	var kArr [26]int
	for i, v := range keyboard {
		kArr[v-'a'] = i
	}
	var pref int
	var count int
	for _, v := range word {
		if kArr[v-'a'] > pref {
			count += kArr[v-'a'] - pref
		} else {
			count += pref - kArr[v-'a']
		}
		pref = kArr[v-'a']
	}
	fmt.Println(count)
	return count
}
