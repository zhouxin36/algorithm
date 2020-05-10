package main

import "fmt"

func main() {
	fmt.Println(countLargestGroup(13) == 4)
	fmt.Println(countLargestGroup(2) == 2)
	fmt.Println(countLargestGroup(15) == 6)
	fmt.Println(countLargestGroup(24) == 5)
	fmt.Println(countLargestGroup(46) == 6)
}
func countLargestGroup(n int) int {
	m := [37]int{}
	var maxCount, res int
	for i := 1; i <= n; i++ {
		//get the sum of digits for current number
		num := i
		var sum int
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		//increment freq counter
		m[sum]++
		//check if this is our new max = start counting from 1, if it next occurance of current max
		//increment count
		if maxCount < m[sum] {
			maxCount = m[sum]
			res = 1
		} else if maxCount == m[sum] {
			res++
		}
	}
	return res
}
