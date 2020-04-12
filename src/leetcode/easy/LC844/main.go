package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
	fmt.Println(!backspaceCompare("a#c", "b"))
}
func backspaceCompare(S string, T string) bool {
	var a []int32
	var b []int32
	for _, s := range S {
		if s == '#' {
			if len(a) != 0 {
				a = a[:len(a)-1]
			}
		} else {
			a = append(a, s)
		}
	}
	for _, t := range T {
		if t == '#' {
			if len(b) != 0 {
				b = b[:len(b)-1]
			}
		} else {
			b = append(b, t)
		}
	}
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
