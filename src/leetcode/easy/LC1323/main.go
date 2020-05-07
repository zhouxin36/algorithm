package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximum69Number(9669) == 9969)
	fmt.Println(maximum69Number(9996) == 9999)
	fmt.Println(maximum69Number(9999) == 9999)
}
func maximum69Number(num int) int {
	b := []byte(strconv.Itoa(num))
	for i, v := range b {
		if v == '6' {
			b[i] = '9'
			break
		}
	}
	res, _ := strconv.Atoi(string(b))
	return res
}
