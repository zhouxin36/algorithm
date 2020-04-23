package main

import "fmt"

func main() {
	//true,false,false
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	//false,false,false
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))
	//true,false,false,false,true,false
	fmt.Println(prefixesDivBy5([]int{0, 1, 1, 1, 1, 1}))
	//false,false,false,false,false
	fmt.Println(prefixesDivBy5([]int{1, 1, 1, 0, 1}))
}
func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	num := 0
	for index, a := range A {
		num = num<<1 | a
		num %= 5
		if num == 0 {
			res[index] = true
		}
	}
	return res
}
