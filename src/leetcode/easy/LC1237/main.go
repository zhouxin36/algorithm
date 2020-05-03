package main

func main() {

}

func findSolution(customFunction func(int, int) int, z int) [][]int {
	x := 1
	y := 1000
	var res [][]int
	for x <= 1000 && y >= 1 {
		if customFunction(x, y) > z {
			y--
		} else if customFunction(x, y) < z {
			x++
		} else {
			res = append(res, []int{x, y})
			x++
			y--
		}
	}
	return res
}
