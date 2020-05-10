package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{2, 3, 4}, 3))
	fmt.Println(buildArray([]int{1, 3}, 3))
}
func buildArray(target []int, n int) []string {
	var res []string
	index := 1
	for _, v := range target {
		if v == index {
			res = append(res, "Push")
		} else {
			for ; index < v; index++ {
				res = append(res, "Push")
				res = append(res, "Pop")
			}
			res = append(res, "Push")
		}
		index++
	}
	return res
}
