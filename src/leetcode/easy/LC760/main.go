package main

import "fmt"

func main() {
	//[]int{1, 4, 3, 2, 0]}
	fmt.Println(anagramMappings([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28}))
}
func anagramMappings(a []int, b []int) []int {
	m := map[int]int{}
	for index, n := range b {
		m[n] = index
	}
	var result []int
	for _, n := range a {
		result = append(result, m[n])
	}
	return result
}
