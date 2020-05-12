package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindromeSubseq("aaa") == 3)
	fmt.Println(longestPalindromeSubseq("bbbab") == 4)
	fmt.Println(longestPalindromeSubseq("cbbd") == 2)
}
func longestPalindromeSubseq(s string) int {
	length := len(s)
	if length < 2 {
		return length
	}
	res := 1
	m := make([]int, length)
	n := make([]int, length)
	p := make([]int, length)
	gap := 0
	for ; gap < length; gap++ {
		for i := gap; i < length; i++ {
			if s[i] == s[i-gap] {
				if gap < 3 {
					p[i-gap] = gap + 1
				} else {
					p[i-gap] = m[i-gap+1] + 2
				}
			} else {
				if n[i-gap] > n[i-gap+1] {
					p[i-gap] = n[i-gap]
				} else {
					p[i-gap] = n[i-gap+1]
				}
			}
			if res < p[i-gap] {
				res = p[i-gap]
			}
		}
		m, n, p = n, p, m
	}
	return res
}

func longestPalindromeSubseq1(s string) int {
	length := len(s)
	if length < 2 {
		return length
	}
	var res int
	arr := make([][]int, length)
	for i := 0; i < length; i++ {
		arr[i] = make([]int, length-i)
		arr[i][0] = 1
	}
	gap := 1
	for ; gap < length; gap++ {
		for i := gap; i < length; i++ {
			if s[i] == s[i-gap] {
				if gap == 1 {
					arr[i-gap][gap] = 2
				} else {
					arr[i-gap][gap] = arr[i-gap+1][gap-2] + 2
				}
			} else {
				if arr[i-gap][gap-1] > arr[i-gap+1][gap-1] {
					arr[i-gap][gap] = arr[i-gap][gap-1]
				} else {
					arr[i-gap][gap] = arr[i-gap+1][gap-1]
				}
			}
			if res < arr[i-gap][gap] {
				res = arr[i-gap][gap]
			}
		}
	}
	return res
}
