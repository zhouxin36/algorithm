package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindromeSubseq1("aaa") == 3)
	fmt.Println(longestPalindromeSubseq1("bbbab") == 4)
	fmt.Println(longestPalindromeSubseq1("cbbd") == 2)
}

/**
1 2 3 ......n
1 2 3 ...
......
1

当列为n时，取得最大值
*/
func longestPalindromeSubseq(s string) int {
	length := len(s)
	if length < 2 {
		return length
	}
	m := make([]int, length)
	n := make([]int, length)
	p := make([]int, length)
	for gap := 0; gap < length; gap++ {
		for i := 0; i < length-gap; i++ {
			if s[i+gap] == s[i] {
				if gap < 3 {
					// gap为0时，子串为1；为1时，子串为2
					p[i] = gap + 1
				} else {
					// i+1为i缩短一个字符，m与p差两个gap；表示除去头和尾的子串
					p[i] = m[i+1] + 2
				}
			} else {
				if n[i] > n[i+1] {
					p[i] = n[i]
				} else {
					p[i] = n[i+1]
				}
			}
		}
		m, n, p = n, p, m
	}
	// 最长子序列长度为整串（gap最大时）
	return n[0]
}

func longestPalindromeSubseq1(s string) int {
	length := len(s)
	if length < 2 {
		return length
	}
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
		}
	}
	// 最长子序列长度为整串（gap最大时）
	return arr[0][length-1]
}
