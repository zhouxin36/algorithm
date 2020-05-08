package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14) == 6)
	fmt.Println(numberOfSteps(8) == 4)
	fmt.Println(numberOfSteps(123) == 12)
}
func numberOfSteps(num int) int {
	var res int
	for num != 0 {
		if num&1 == 0 {
			res++
			num >>= 1
		} else {
			if num == 1 {
				res++
				break
			} else {
				res += 2
				num >>= 1
			}
		}
	}
	return res
}
