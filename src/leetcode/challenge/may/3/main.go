package main

import "fmt"

func main() {
	fmt.Println(!canConstruct("a", "b"))
	fmt.Println(!canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
func canConstruct(ransomNote string, magazine string) bool {
	m := map[int32]int{}
	for _, v := range ransomNote {
		m[v]++
	}
	for _, v := range magazine {
		if len(m) == 0 {
			return true
		}
		if m[v] > 1 {
			m[v]--
		} else if m[v] == 1 {
			delete(m, v)
		}
	}
	return len(m) == 0
}
