package main

import "fmt"

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}
func processQueries(queries []int, m int) []int {
	ma := make([]int, m+1)
	result := make([]int, len(queries))
	for i := 0; i <= m; i++ {
		ma[i] = i - 1
	}
	for index, q := range queries {
		result[index] = ma[q]
		count := ma[q]
		for i := 1; i <= m; i++ {
			if ma[i] < ma[q] {
				ma[i]++
				count--
			}
			if count == 0 {
				break
			}
		}
		ma[q] = 0
	}
	return result
}
