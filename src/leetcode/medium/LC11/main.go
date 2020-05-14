package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 2, 4, 3}) == 4)
	fmt.Println(maxArea([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}) == 25)
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
	fmt.Println(maxArea([]int{1, 1}) == 1)
}
func maxArea(height []int) int {
	start, end := 0, len(height)-1
	var res int
	for start < end {
		var v int
		if height[start] < height[end] {
			v = height[start] * (end - start)
			start++
		} else {
			v = height[end] * (end - start)
			end--
		}
		if res < v {
			res = v
		}
	}
	return res
}
func maxArea2(height []int) int {
	var res int
	var min, max []int
	for i := 0; i < len(height); i++ {
		if len(max) == 0 || height[max[len(max)-1]] < height[i] {
			max = append(max, i)
		}
	}
	for i := len(height) - 1; i >= 0; i-- {
		if len(min) == 0 || height[min[len(min)-1]] < height[i] {
			min = append(min, i)
		}
	}
	for i := 0; i < len(max); i++ {
		for j := 0; j < len(min) && max[i] < min[j]; j++ {
			var h int
			if height[min[j]] < height[max[i]] {
				h = height[min[j]]
			} else {
				h = height[max[i]]
			}
			v := h * (min[j] - max[i])
			if v > res {
				res = v
			}
		}
	}
	return res
}

func maxArea1(height []int) int {
	var res int
	for i, _ := range height {
		for j := i + 1; j < len(height); j++ {
			var h int
			if height[i] > height[j] {
				h = height[j]
			} else {
				h = height[i]
			}
			v := h * (j - i)
			if res < v {
				res = v
			}
		}
	}
	return res
}
