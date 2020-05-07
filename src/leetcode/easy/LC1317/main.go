package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1,1
	fmt.Println(getNoZeroIntegers(2))
	//2,9
	fmt.Println(getNoZeroIntegers(11))
	//1,9999
	fmt.Println(getNoZeroIntegers(10000))
	//1,68
	fmt.Println(getNoZeroIntegers(69))
	//11,999
	fmt.Println(getNoZeroIntegers(1010))
}
func getNoZeroIntegers(n int) []int {
	i := 1
	for {
		if !strings.Contains(strconv.Itoa(i)+strconv.Itoa(n-i), "0") {
			break
		} else {
			i++
		}
	}
	return []int{i, n - i}
}
