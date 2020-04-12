package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	fmt.Println(lemonadeChange([]int{5, 5, 10}))
	fmt.Println(!lemonadeChange([]int{10, 10}))
	fmt.Println(!lemonadeChange([]int{5, 5, 10, 10, 20}))
}
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		switch b {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		case 20:
			if ten > 0 && five > 0 {
				five--
				ten--

			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
