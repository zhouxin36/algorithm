package main

import "fmt"

func main() {
	fmt.Println(dietPlan([]int{1, 2, 3, 4, 5}, 1, 3, 3) == 0)
	fmt.Println(dietPlan([]int{3, 2}, 2, 0, 1) == 1)
	fmt.Println(dietPlan([]int{6, 5, 0, 0}, 2, 1, 5) == 0)
}
func dietPlan(arr []int, k, lower, upper int) int {
	var sum int
	var res int
	for _, v := range arr[:k] {
		sum += v
	}
	if sum > upper {
		res++
	} else if sum < lower {
		res--
	}
	for i := k; i < len(arr); i++ {
		sum += arr[i]
		sum -= arr[i-k]
		if sum > upper {
			res++
		} else if sum < lower {
			res--
		}
	}
	return res
}
