package main

import "fmt"

func main() {
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7}))
}
func countTriplets(arr []int) int {
	var res int
	for i := 0; i < len(arr); i++ {
		a := arr[i]
		for j := i + 1; j < len(arr); j++ {
			b := 0
			for k := j; k < len(arr); k++ {
				b ^= arr[k]
				if a == b {
					res++
				}
			}
			a ^= arr[j]
		}
	}
	return res
}
