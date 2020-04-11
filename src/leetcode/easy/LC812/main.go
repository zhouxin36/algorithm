package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}))
}
func largestTriangleArea(points [][]int) float64 {
	var ans float64
	length := len(points)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				v := area(points[i], points[j], points[k])
				if v > ans {
					ans = v
				}
			}
		}
	}
	return ans
}

func area(P []int, Q []int, R []int) float64 {
	return 0.5 * math.Abs(float64(P[0]*Q[1]+Q[0]*R[1]+R[0]*P[1]-P[1]*Q[0]-Q[1]*R[0]-R[1]*P[0]))
}
