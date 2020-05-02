package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb") == 3)
	fmt.Println(numJewelsInStones("z", "ZZ") == 0)
}
func numJewelsInStones(J string, S string) int {
	j := map[byte]bool{}
	for _, v := range J {
		j[byte(v)] = true
	}
	var count int
	for _, v := range S {
		if j[byte(v)] {
			count++
		}
	}
	return count
}
