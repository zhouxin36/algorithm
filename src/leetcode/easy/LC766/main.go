package main

import (
	"fmt"
)

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{41, 45}, {81, 41}, {73, 81}, {47, 73}, {76, 47}, {79, 76}}))
	fmt.Println(isToeplitzMatrix([][]int{{83}, {64}, {57}}))
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2}}))

	fmt.Println(!isToeplitzMatrix([][]int{{1, 2}, {1, 2}}))
}
func isToeplitzMatrix(matrix [][]int) bool {
	remark := make([]int, len(matrix[0]))
	for index, v := range matrix[0] {
		remark[index] = v
	}
	length := len(remark)
	index := length - 1
	for _, v := range matrix[1:] {
		for j, m := range v {
			if j == 0 {
				remark[index] = m
			} else {
				if m != remark[(index+j)%length] {
					return false
				}
			}
		}
		if index > 0 {
			index--
		} else {
			index = length - 1
		}
	}
	return true
}
