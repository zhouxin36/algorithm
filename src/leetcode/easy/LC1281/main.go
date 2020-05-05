package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234) == 15)
	fmt.Println(subtractProductAndSum(4421) == 21)
}
func subtractProductAndSum(n int) int {
	var sum int
	product := 1
	for n != 0 {
		v := n % 10
		n /= 10
		sum += v
		product *= v
	}
	return product - sum
}
