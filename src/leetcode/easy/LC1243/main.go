package main

import "fmt"

func main() {
	//6,3,3,4
	fmt.Println(transformArray([]int{6, 2, 3, 4}))
	//1,4,4,4,4,5
	fmt.Println(transformArray([]int{1, 6, 3, 4, 3, 5}))
}
func transformArray(arr []int) []int {
	for {
		var flag bool
		for i := 1; i < len(arr)-2; i++ {
			if arr[i-1] < arr[i] && arr[i+1] < arr[i] {
				arr[i]--
				flag = true
			} else if arr[i-1] > arr[i] && arr[i+1] > arr[i] {
				arr[i]++
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}
