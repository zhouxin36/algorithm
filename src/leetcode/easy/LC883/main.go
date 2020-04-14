package main

import "fmt"

func main() {
	fmt.Println(projectionArea([][]int{{2}}) == 5)
	fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}}) == 17)
	fmt.Println(projectionArea([][]int{{1, 0}, {0, 2}}) == 8)
	fmt.Println(projectionArea([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}) == 14)
	fmt.Println(projectionArea([][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}) == 21)
}

func projectionArea(grid [][]int) int {
	var x, y, z int
	for _, g := range grid {
		max := 0
		for _, gg := range g {
			if gg > 0 {
				z++
				if max < gg {
					max = gg
				}
			}
		}
		x += max
	}
	for j := 0; j < len(grid[0]); j++ {
		max := 0
		for i := 0; i < len(grid); i++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		y += max
	}
	return x + y + z
}
