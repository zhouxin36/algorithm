package main

import "fmt"

func main() {
	fmt.Println(reformat("j"))
	fmt.Println(reformat("j1"))
	fmt.Println(reformat("j12"))
	fmt.Println(reformat("1"))
	fmt.Println(reformat("ab123"))
	fmt.Println(reformat("abce123"))
	fmt.Println(reformat("a0b1c2"))
	fmt.Println(reformat("leetcode"))
	fmt.Println(reformat("1229857369"))
	fmt.Println(reformat("covid2019"))
}
func reformat(s string) string {
	if len(s) < 2 {
		return s
	}
	var result []int32
	var a []int32
	var b []int32
	for _, v := range s {
		if v >= '0' && v <= '9' {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	if a == nil || b == nil {
		return ""
	}
	if len(a)+1 == len(b) || len(a) == len(b) {
		for i, av := range a {
			result = append(result, b[i])
			result = append(result, av)
		}
		if len(a) != len(b) {
			result = append(result, b[len(b)-1])
		}
		return string(result)
	} else if len(a) == len(b)+1 {
		for i, bv := range b {
			result = append(result, a[i])
			result = append(result, bv)
		}
		result = append(result, a[len(a)-1])
		return string(result)
	} else {
		return ""
	}
}
