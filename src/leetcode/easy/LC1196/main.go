package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(apple([]int{100, 200, 150, 1000}) == 4)
	fmt.Println(apple([]int{900, 950, 800, 1000, 700, 800}) == 5)
}
func apple(apples []int) int {
	sort.Ints(apples)
	var sum int

	for i := 0; i < len(apples); i++ {
		if sum < 5000 {
			sum += apples[i]
		} else {
			return i
		}
	}
	if sum > 5000 {
		return len(apples) - 1
	} else {
		return len(apples)
	}

}
