package main

import "fmt"

func main() {
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	fmt.Println([][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}})
}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	old := image[sr][sc]
	image[sr][sc] = newColor
	if sr > 0 && image[sr-1][sc] == old {
		floodFill(image, sr-1, sc, newColor)
	}
	if sc > 0 && image[sr][sc-1] == old {
		floodFill(image, sr, sc-1, newColor)
	}
	if sr < len(image)-1 && image[sr+1][sc] == old {
		floodFill(image, sr+1, sc, newColor)
	}
	if sc < len(image[sr])-1 && image[sr][sc+1] == old {
		floodFill(image, sr, sc+1, newColor)
	}
	return image
}
