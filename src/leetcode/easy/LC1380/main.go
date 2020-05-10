package main

import "fmt"

func main() {
	//12
	fmt.Println(luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}))
	//15
	fmt.Println(luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))
	//7
	fmt.Println(luckyNumbers([][]int{{7, 8}, {1, 2}}))
}
func luckyNumbers(matrix [][]int) []int {
	var res []int
	for i, row := range matrix {
		min := 0
		for ri := 1; ri < len(row); ri++ {
			if row[min] > row[ri] {
				min = ri
			}
		}
		var flag bool
		for j := 0; j < len(matrix); j++ {
			if matrix[j][min] > matrix[i][min] {
				flag = true
				break
			}
		}
		if !flag {
			res = append(res, matrix[i][min])
		}
	}
	return res
}
