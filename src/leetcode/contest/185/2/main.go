package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//[["Table","Beef Burrito","Ceviche","Fried Chicken","Water"],["3","0","2","1","0"],["5","0","1","0","1"],["10","1","0","0","0"]]
	fmt.Println(displayTable([][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}}))
}
func displayTable(orders [][]string) [][]string {
	sort.Slice(orders, func(i, j int) bool {
		return strings.Compare(orders[i][2], orders[j][2]) != 1
	})
	var name []string
	nm := map[string]int{}
	var result [][]string
	var nums [][]int
	name = append(name, "Table")
	for _, iv := range orders {
		if _, ok := nm[iv[2]]; !ok {
			nm[iv[2]] = len(name)
			name = append(name, iv[2])
		}
	}
	sort.Slice(orders, func(i, j int) bool {
		iv, _ := strconv.Atoi(orders[i][1])
		jv, _ := strconv.Atoi(orders[j][1])
		return iv <= jv
	})
	sm := map[string]int{}
	for _, v := range orders {
		sv, ok := sm[v[1]]
		if !ok {
			num := make([]int, len(name))
			num[nm[v[2]]]++
			sm[v[1]] = len(nums)
			va, _ := strconv.Atoi(v[1])
			num[0] = va
			nums = append(nums, num)
		} else {
			nums[sv][nm[v[2]]]++

		}
	}
	result = append(result, name)
	for _, n := range nums {
		res := make([]string, len(name))
		for j, jv := range n {
			res[j] = strconv.Itoa(jv)
		}
		result = append(result, res)
	}
	return result
}
