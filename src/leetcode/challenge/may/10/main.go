package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(!isPerfectSquare(14))
}
func isPerfectSquare(num int) bool {
	if num == 0 {
		return false
	} else if num == 1 {
		return true
	}
	pref, suffix := num, num
	for pref*pref > num {
		suffix = pref
		pref >>= 1
	}
	for i := pref; i < suffix; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
