package main

import "fmt"

func main() {
	//18,6,6,6,1,-1
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}
func replaceElements(arr []int) []int {
	right := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = right
		if right < tmp {
			right = tmp
		}
	}
	return arr
}
