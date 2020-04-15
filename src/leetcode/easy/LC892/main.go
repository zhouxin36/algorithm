package main

import "fmt"

func main() {
	fmt.Println(surfaceArea([][]int{{2}}) == 10)
	fmt.Println(surfaceArea([][]int{{1, 2}, {3, 4}}) == 34)
	fmt.Println(surfaceArea([][]int{{1, 0}, {0, 2}}) == 16)
	fmt.Println(surfaceArea([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}) == 32)
	fmt.Println(surfaceArea([][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}) == 46)
}
func surfaceArea(grid [][]int) int {
	var total, cover int
	for in, gr := range grid {
		for i, g := range gr {
			if g == 0 {
				continue
			} else {
				cover += g - 1
			}
			total += g
			if in > 0 {
				v := grid[in-1][i]
				if v > g {
					cover += g
				} else {
					cover += v
				}
			}
			if i > 0 {
				v := grid[in][i-1]
				if v > g {
					cover += g
				} else {
					cover += v
				}
			}
		}
	}
	return 6*total - 2*cover
}
