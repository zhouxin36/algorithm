package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(!uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
func uniqueOccurrences(arr []int) bool {
	count := map[int]int{}
	for _, v := range arr {
		count[v]++
	}
	hit := map[int]bool{}
	for _, v := range count {
		if hit[v] {
			return false
		} else {
			hit[v] = true
		}
	}
	return true
}
