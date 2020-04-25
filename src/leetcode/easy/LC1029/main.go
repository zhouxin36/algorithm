package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}) == 110)
	fmt.Println(twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}) == 1859)
}
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0]-costs[i][1])-(costs[j][0]-costs[j][1]) < 0
	})
	var sum int
	var half = len(costs) >> 1
	for index, c := range costs {
		if index < half {
			sum += c[0]
		} else {
			sum += c[1]
		}
	}
	return sum
}
