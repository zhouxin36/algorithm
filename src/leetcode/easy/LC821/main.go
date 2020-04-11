package main

import (
	"fmt"
	"math"
)

func main() {
	//[3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
func shortestToChar(S string, C byte) []int {
	arr := make([]int, len(S))
	pref := math.MinInt32
	for index, s := range S {
		if byte(s) == C {
			pref = index
		}
		arr[index] = index - pref
	}
	suf := math.MaxInt32
	for i := len(S) - 1; i >= 0; i-- {
		if byte(S[i]) == C {
			suf = i
		}
		if arr[i] > suf-i {
			arr[i] = suf - i
		}
	}
	return arr
}
