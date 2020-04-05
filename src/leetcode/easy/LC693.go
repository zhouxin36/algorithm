package main

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(!hasAlternatingBits(7))
	fmt.Println(!hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(10))
}
func hasAlternatingBits(n int) bool {
	return (n&(n>>1)) == 0 && (n&(n>>2)) == (n>>2)
}
func hasAlternatingBits1(n int) bool {
	i := uint(n)
	if i%2 == 0 {
		i >>= 1
	}
	for i != 0 {
		if i%2 == 1 {
			i >>= 1
		} else {
			return false
		}
		if i%2 == 0 {
			i >>= 1
		} else {
			return false
		}
	}
	return true
}
