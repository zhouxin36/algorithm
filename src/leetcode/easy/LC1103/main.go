package main

import (
	"fmt"
)

func main() {
	//[1,2,3,1]
	fmt.Println(distributeCandies(7, 4))
	//[5,2,3]
	fmt.Println(distributeCandies(10, 3))
}
func distributeCandies(candies int, numPeople int) []int {
	res := make([]int, numPeople)
	for i := 1; candies > 0; i++ {
		if candies >= i {
			res[(i-1)%numPeople] += i
			candies -= i
		} else {
			res[(i-1)%numPeople] += candies
			break
		}
	}
	return res
}
