package main

import "fmt"

func main() {

	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}) == 3)
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}) == 1)
}
func numEquivDominoPairs1(dominoes [][]int) int {
	if len(dominoes) < 2 {
		return 0
	}
	var count int
	for i := 0; i < len(dominoes); i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) || (dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0]) {
				count++
			}
		}
	}
	return count
}
func numEquivDominoPairs2(dominoes [][]int) int {
	if len(dominoes) < 2 {
		return 0
	}
	var count int
	record := map[int][]int{}
	for i, v := range dominoes {
		record[v[0]+v[1]] = append(record[v[0]+v[1]], i)
	}
	for _, v := range record {
		if len(v) < 2 {
			continue
		}
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				if (dominoes[v[i]][0] == dominoes[v[j]][0] && dominoes[v[i]][1] == dominoes[v[j]][1]) || (dominoes[v[i]][0] == dominoes[v[j]][1] && dominoes[v[i]][1] == dominoes[v[j]][0]) {
					count++
				}
			}
		}
	}
	return count
}
func numEquivDominoPairs(dominoes [][]int) int {
	record := map[int]int{}
	for _, v := range dominoes {
		var a int
		if v[0] > v[1] {
			a = v[0]*10 + v[1]
		} else {
			a = v[1]*10 + v[0]
		}
		record[a]++
	}
	var res int
	for _, v := range record {
		res += v * (v - 1)
	}
	return res >> 1
}
