package main

import "fmt"

func main() {
	fmt.Println(minTime(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 4}, {2, 5}, {2, 6}, {4, 7}},
		[]bool{true, true, false, true, false, true, true, false}))
	fmt.Println(minTime(4, [][]int{{0, 1}, {1, 2}, {0, 3}}, []bool{true, true, true, true}))
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
		[]bool{false, false, true, false, true, true, false}))
}
func minTime(n int, edges [][]int, hasApple []bool) int {

	var count int
	for _, v := range hasApple {
		if v {
			count++
		}
	}
	if count == 0 {
		return 0
	}
	ma := map[int][]int{}
	for _, v := range edges {
		ma[v[0]] = append(ma[v[0]], v[1])
	}
	var steps [][2]int
	m := map[[2]int]bool{}
	tree(&steps, &m, &ma, &hasApple, 0)
	fmt.Println(m)
	return len(m) << 1
}
func tree(steps *[][2]int, m *map[[2]int]bool, ma *map[int][]int, hasApple *[]bool, index int) {
	if (*hasApple)[index] {
		for _, v := range *steps {
			(*m)[v] = true
		}
	}
	if len((*ma)[index]) == 0 {
		return
	}
	for _, v := range (*ma)[index] {
		*steps = append(*steps, [2]int{index, v})
		tree(steps, m, ma, hasApple, v)
		*steps = (*steps)[:len(*steps)-1]
	}
}
