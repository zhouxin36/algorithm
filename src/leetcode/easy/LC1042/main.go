package main

import "fmt"

func main() {
	//1,2,3
	fmt.Println(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	//1,2,1,2
	fmt.Println(gardenNoAdj(4, [][]int{{1, 2}, {3, 4}}))
	//1,2,3,4
	fmt.Println(gardenNoAdj(4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}))
}
func gardenNoAdj(N int, paths [][]int) []int {
	arr := make([][]int, N)
	for _, ar := range paths {
		arr[ar[0]-1] = append(arr[ar[0]-1], ar[1]-1)
		arr[ar[1]-1] = append(arr[ar[1]-1], ar[0]-1)
	}
	res := make([]int, N)
	flower := make([]bool, 4)
	for index, a := range arr {
		for i := range flower {
			flower[i] = false
		}
		for _, v := range a {
			if res[v] > 0 {
				flower[res[v]-1] = true
			}
		}
		for i := 0; i < 4; i++ {
			if !flower[i] {
				res[index] = i + 1
				break
			}
		}
	}
	return res
}
