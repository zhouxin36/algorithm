package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(armstrongNumber(153))
	fmt.Println(!armstrongNumber(123))
}
func armstrongNumber(N int) bool {
	var res int
	ns := strconv.Itoa(N)
	for i := 0; i < len(ns); i++ {
		res += pow(int(ns[i]-'0'), len(ns))
	}
	if res == N {
		return true
	} else {
		return false
	}
}

func pow(x, y int) int {
	if y == 0 {
		return 1
	} else if y == 1 {
		return x
	} else if y == 2 {
		return x * x
	}
	if y&1 == 0 {
		v := pow(x, y>>1)
		return v * v
	} else {
		v := pow(x, y>>1)
		return v * v * x
	}
}
