package main

import "fmt"

func main() {
	fmt.Println(isBoomerang([][]int{{0, 1}, {2, 0}, {1, 1}}))
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}}))
	fmt.Println(!isBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}}))
}
func isBoomerang2(points [][]int) bool {
	if points[0][0] == points[1][0] && points[0][1] == points[1][1] {
		return false
	} else if points[2][0] == points[1][0] && points[2][1] == points[1][1] {
		return false
	} else if points[2][0] == points[0][0] && points[2][1] == points[0][1] {
		return false
	}
	if points[0][0] == points[1][0] {
		if points[0][0] == points[2][0] {
			return false
		} else {
			return true
		}
	}
	a := float64(points[0][1]-points[1][1]) / float64(points[0][0]-points[1][0])
	b := float64(points[0][1]) - a*float64(points[0][0])
	return a*float64(points[2][0])+b != float64(points[2][1])
}

/**
正弦
*/
func isBoomerang(points [][]int) bool {
	dx1 := points[1][0] - points[0][0]
	dx2 := points[2][0] - points[1][0]
	dy1 := points[1][1] - points[0][1]
	dy2 := points[2][1] - points[1][1]

	return dy1*dx2 != dy2*dx1
}
