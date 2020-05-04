package main

import "fmt"

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}) == 7)
	fmt.Println(minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}) == 5)
}
func minTimeToVisitAllPoints(points [][]int) int {
	var count int
	for i := 1; i < len(points); i++ {
		y := abs(points[i][0] - points[i-1][0])
		x := abs(points[i][1] - points[i-1][1])
		if y > x {
			count += y
		} else {
			count += x
		}
	}
	return count
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
