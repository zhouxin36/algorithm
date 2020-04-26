package main

import (
	"fmt"
	"strings"
)

func main() {
	//[[3,7],[9,13],[10,17]]
	fmt.Println(indexPairs("thestoryofleetcodeandme", []string{"story", "fleet", "leetcode"}))
	//[[0,1],[0,2],[2,3],[2,4]]
	fmt.Println(indexPairs("ababa", []string{"aba", "ab"}))
}
func indexPairs(str string, strs []string) [][]int {
	var res [][]int
	for _, v := range strs {
		var index int
		tmp := str
		for true {
			index = strings.Index(tmp, v)
			if index != -1 {
				res = append(res, []int{index, index + len(v) - 1})
				v := []byte(tmp)
				v[index] = '0'
				tmp = string(v)
			} else {
				break
			}
		}
	}
	return res
}
