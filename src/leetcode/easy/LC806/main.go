package main

import "fmt"

func main() {
	//[3, 60]
	fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}
func numberOfLines(widths []int, S string) []int {
	var result = make([]int, 2)
	result[0] = 1
	for _, s := range S {
		i := widths[s-'a']
		if result[1]+i > 100 {
			result[0]++
			result[1] = i
		} else {
			result[1] += i
		}
	}
	return result
}
