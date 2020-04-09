package main

import "fmt"

func main() {
	fmt.Println(countPrimeSetBits(6, 10) == 4)
	fmt.Println(countPrimeSetBits(10, 15) == 5)
}
func countPrimeSetBits(L int, R int) int {
	count := 0
	for i := L; i <= R; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}
func isPrime(num int) bool {
	count := 0
	for num != 0 {
		count++
		num &= num - 1
	}
	switch count {
	case 2, 3, 5, 7, 11, 13, 17, 19:
		return true
	default:
		return false
	}
}
