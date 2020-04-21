package main

import "fmt"

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}) == 2)
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}) == 3)
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}) == -1)
	fmt.Println(findJudge(3, [][]int{{1, 2}, {2, 3}}) == -1)
	fmt.Println(findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}) == 3)
}
func findJudge(N int, trust [][]int) int {
	nArr := make([]int, N)
	for _, t := range trust {
		nArr[t[0]-1] = -1
		if nArr[t[1]-1] != -1 {
			nArr[t[1]-1]++
		}
	}
	for index, n := range nArr {
		if n == len(nArr)-1 {
			return index + 1
		}
	}
	return -1
}
