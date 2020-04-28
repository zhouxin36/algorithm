package main

import "fmt"

func main() {
	fmt.Println(tribonacci(4) == 4)
	fmt.Println(tribonacci(25) == 1389537)
}
func tribonacci3(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	default:
		return tribonacci3(n-3) + tribonacci3(n-2) + tribonacci3(n-1)
	}
}
func tribonacci2(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	case 3:
		return 2
	}
	v := make([]int, n+1)
	v[0] = 0
	v[1] = 1
	v[2] = 1
	v[3] = 2
	return doTribonacci2(n, &v)
}
func doTribonacci2(n int, arr *[]int) int {
	if (*arr)[n] != 0 {
		return (*arr)[n]
	}
	v := doTribonacci2(n-3, arr) + doTribonacci2(n-2, arr) + doTribonacci2(n-1, arr)
	(*arr)[n] = v
	return v
}
func tribonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	case 3:
		return 2
	}
	remark := [3]int{2, 1, 1}
	for i := 4; i <= n; i++ {
		index := i % 3
		remark[index] = remark[0] + remark[1] + remark[2]
	}
	return remark[n%3]
}
