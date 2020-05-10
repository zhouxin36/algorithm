package main

import "fmt"

func main() {
	//true,true,true,false,true
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	//true,false,false,false,false
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	//true,false,true
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	res := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= max {
			res[i] = true
		}
	}
	return res
}
