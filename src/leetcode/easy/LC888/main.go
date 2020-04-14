package main

import "fmt"

func main() {
	//1,2
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))
	//1,2
	fmt.Println(fairCandySwap([]int{1, 2}, []int{2, 3}))
	//2,3
	fmt.Println(fairCandySwap([]int{2}, []int{1, 3}))
	//5,4
	fmt.Println(fairCandySwap([]int{1, 2, 5}, []int{2, 4}))
}
func fairCandySwap(A []int, B []int) []int {
	var sumA, sumB int
	m := map[int]bool{}
	for _, a := range A {
		m[a] = true
		sumA += a
	}
	for _, b := range B {
		sumB += b
	}
	dif := (sumA - sumB) / 2
	for _, b := range B {
		if m[b+dif] {
			return []int{b + dif, b}
		}
	}
	return nil
}
