package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"36", "28", "70", "65", "C", "+", "33", "-46", "84", "C"}) == 219)
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}) == 30)
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}) == 27)
}

/**
栈
*/
//func calPoints1(ops []string) int {
//	var stack []int
//	for _, op := range ops {
//		switch op {
//		case "+":
//			v := stack[len(stack)-1] + stack[len(stack)-2]
//			stack = append(stack, v)
//		case "C":
//			stack = stack[:len(stack)-1]
//		case "D":
//			stack = append(stack, 2 * stack[len(stack)-1])
//		default:
//			v, _ := strconv.Atoi(op)
//			stack = append(stack, v)
//		}
//	}
//	var sum int
//	for _, n := range stack {
//		sum += n
//	}
//	return sum
//}

/**
数组
*/
func calPoints(ops []string) int {
	length := len(ops)
	var arr = make([]int, length)
	size := 0
	for i := 0; i < length; i++ {
		switch ops[i][0] {
		case 'C':
			size--
			break
		case 'D':
			arr[size] = arr[size-1] * 2
			size++
			break
		case '+':
			arr[size] = arr[size-1] + arr[size-2]
			size++
			break
		default:
			arr[size], _ = strconv.Atoi(ops[i])
			size++
		}
	}
	var result int
	for i := 0; i < size; i++ {
		result += arr[i]
	}
	return result
}
