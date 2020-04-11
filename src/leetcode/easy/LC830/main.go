package main

import "fmt"

func main() {
	//[[3,6]]
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
	//[[3,5],[6,9],[12,14]]
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
}
func largeGroupPositions(S string) [][]int {
	var result [][]int
	pref := '1'
	count := 1
	for index, s := range S {
		if pref == s {
			count++
		} else if count >= 3 {
			result = append(result, []int{index - count, index - 1})
			count = 1
		} else {
			count = 1
		}
		pref = s
	}
	if count >= 3 {
		result = append(result, []int{len(S) - count, len(S) - 1})
	}
	return result
}
