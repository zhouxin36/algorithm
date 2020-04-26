package main

import "fmt"

func main() {
	var arr []int
	arr = []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
	arr = []int{1, 2, 3}
	duplicateZeros(arr)
	fmt.Println(arr)
}
func duplicateZeros1(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			i++
		}
	}
}
func duplicateZeros(arr []int) {
	a := make([]int, len(arr))
	var index int
	for i := 0; i < len(arr) && index < len(arr); i++ {
		if arr[i] == 0 {
			a[index] = arr[i]
			index++
			if index >= len(arr) {
				break
			}
		}
		a[index] = arr[i]
		index++
	}
	copy(arr, a)
}
