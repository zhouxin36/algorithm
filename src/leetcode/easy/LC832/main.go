package main

import "fmt"

func main() {
	//[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
	//[[1,0,0],[0,1,0],[1,1,1]]
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}
func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		length := len(a)
		for i := 0; i < (length+1)>>1; i++ {
			if a[i] == 0 && a[length-i-1] == 0 {
				a[i] = 1
				a[length-i-1] = 1
			} else if a[i] == 1 && a[length-i-1] == 1 {
				a[i] = 0
				a[length-i-1] = 0
			}
		}
	}
	return A
}
