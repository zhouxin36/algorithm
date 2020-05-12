package main

import "fmt"

func main() {
	//[[2,2,2],[2,2,0],[2,0,1]]
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	return *doFloodFill(&image, sr, sc, image[sr][sc], newColor, len(image), len(image[0]))
}

func doFloodFill(i *[][]int, sr, sc, oldColor, color, row, column int) *[][]int {
	if (*i)[sr][sc] != oldColor {
		return i
	}
	(*i)[sr][sc] = color
	if sr > 0 {
		doFloodFill(i, sr-1, sc, oldColor, color, row, column)
	}
	if sr < row-1 {
		doFloodFill(i, sr+1, sc, oldColor, color, row, column)
	}
	if sc > 0 {
		doFloodFill(i, sr, sc-1, oldColor, color, row, column)
	}
	if sc < column-1 {
		doFloodFill(i, sr, sc+1, oldColor, color, row, column)
	}
	return i
}
