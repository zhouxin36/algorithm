package main

import "fmt"

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3) == 12)
	fmt.Println(maxScore([]int{2, 2, 2}, 2) == 4)
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7) == 55)
	fmt.Println(maxScore([]int{1, 1000, 1}, 1) == 1)
	fmt.Println(maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3) == 202)
}
func maxScore(cardPoints []int, k int) int {
	var i, num int
	for ; i < k; i++ {
		num += cardPoints[i]
	}
	if i >= len(cardPoints) {
		return num
	}
	j, max := len(cardPoints)-1, num
	i--
	for ; i >= 0; i-- {
		num -= cardPoints[i]
		num += cardPoints[j]
		if num > max {
			max = num
		}
		j--
	}
	fmt.Println(max)
	return max
}
