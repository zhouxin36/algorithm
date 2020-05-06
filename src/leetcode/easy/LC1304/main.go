package main

import "fmt"

func main() {
	//-7,-1,1,3,4
	fmt.Println(sumZero(5))
	//-1,0,1
	fmt.Println(sumZero(3))
	//0
	fmt.Println(sumZero(1))
}
func sumZero(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i*2 - n + 1
	}
	return res
}
