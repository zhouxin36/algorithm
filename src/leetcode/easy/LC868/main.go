package main

import "fmt"

func main() {
	fmt.Println(binaryGap(22) == 2)
	fmt.Println(binaryGap(5) == 2)
	fmt.Println(binaryGap(6) == 1)
	fmt.Println(binaryGap(8) == 0)
}
func binaryGap(N int) int {
	var result int
	var pref, cur int
	for N != 0 {
		cur = N & (-N)
		N &= N - 1
		if pref == 0 {
			pref = cur
			continue
		}
		v := cur / pref
		if result < v {
			result = v
		}
		pref = cur
	}
	count := 0
	for result > 1 {
		result >>= 1
		count++
	}
	return count
}
