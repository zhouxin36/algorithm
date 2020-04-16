package main

import "fmt"

func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	fmt.Println(!hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println(!hasGroupsSizeX([]int{1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))
}
func hasGroupsSizeX(deck []int) bool {
	m := map[int]int{}
	for _, de := range deck {
		m[de]++
	}
	pref := 0
	for _, v := range m {
		if v < 2 {
			return false
		}
		if pref == 0 {
			pref = v
		} else {
			if v != pref && GCD(pref, v) < 2 {
				return false
			}
			pref = v
		}
	}
	return true
}
func GCD(a int, b int) int {
	if a < b {
		a ^= b
		b ^= a
		a ^= b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
