package main

import "fmt"

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}) == 3)
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}) == 3)
}
func numPairsDivisibleBy601(time []int) int {
	var count int
	for i := 0; i < len(time); i++ {
		for j := i + 1; j < len(time); j++ {
			if (time[i]+time[j])%60 == 0 {
				count++
			}
		}
	}
	return count
}
func numPairsDivisibleBy60(time []int) int {
	var count int
	var arr [60]int

	for _, t := range time {
		v := t % 60
		count += arr[(60-v)%60]
		arr[v]++
	}
	return count
}
