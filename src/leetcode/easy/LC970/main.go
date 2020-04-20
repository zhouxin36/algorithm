package main

import "fmt"

func main() {
	fmt.Println(powerfulIntegers(2, 1, 10))
	//2,3,4,5,7,9,10
	fmt.Println(powerfulIntegers(2, 3, 10))
	//2,4,6,8,10,14
	fmt.Println(powerfulIntegers(3, 5, 15))
}
func powerfulIntegers(x int, y int, bound int) []int {
	m := map[int]bool{}
	in := 1
	for i := 0; ; i++ {
		if i != 0 {
			in *= x
		}
		if in >= bound {
			break
		}
		jn := 1
		for j := 0; ; j++ {
			if j != 0 {
				jn *= y
			}
			if jn+in > bound {
				break
			}
			m[in+jn] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	var res []int
	for k := range m {
		res = append(res, k)
	}
	return res
}
