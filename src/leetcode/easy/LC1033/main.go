package main

import (
	"fmt"
	"sort"
)

func main() {
	//[1,2]
	fmt.Println(numMovesStones(1, 2, 5))
	//[0,0]
	fmt.Println(numMovesStones(4, 3, 2))
	//[1,2]
	fmt.Println(numMovesStones(3, 5, 1))
}
func numMovesStones(a int, b int, c int) []int {
	ar := []int{a, b, c}
	sort.Ints(ar)
	d1 := ar[1] - ar[0] - 1
	d2 := ar[2] - ar[1] - 1
	switch true {
	case d1 == 0 && d2 == 0:
		return []int{0, 0}
	case d1 == 0 || d1 == 1:
		return []int{1, d1 + d2}
	case d2 == 0 || d2 == 1:
		return []int{1, d1 + d2}
	default:
		return []int{2, d1 + d2}
	}
}
