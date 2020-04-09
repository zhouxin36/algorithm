package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb") == 3)
	fmt.Println(numJewelsInStones("z", "ZZ") == 0)
}
func numJewelsInStones(J string, S string) int {
	m := map[int32]bool{}
	for _, a := range J {
		m[a] = true
	}
	count := 0
	for _, a := range S {
		if m[a] {
			count++
		}
	}
	return count
}
