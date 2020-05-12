package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome4("abcdbbfcba") == "bb")
	fmt.Println(longestPalindrome4("bb") == "bb")
	fmt.Println(longestPalindrome4("abcda") == "a")
	fmt.Println(longestPalindrome4("abcba") == "abcba")
	fmt.Println(longestPalindrome4("babad") == "bab")
	fmt.Println(longestPalindrome4("cbbd") == "bb")
}

/**
Manacher算法
时间复杂度：O(N)
空间复杂度：O(N)
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	bytes := make([]byte, 2*len(s)+1)
	p := make([]int, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		bytes[2*i] = '#'
		bytes[2*i+1] = s[i]
	}
	bytes[len(bytes)-1] = '#'
	var maxRight, center, start int
	for i := 0; i < len(p); i++ {
		if i < maxRight {
			mirror := 2*center - i
			p[i] = maxRight - i
			if p[i] > p[mirror] {
				p[i] = p[mirror]
			}
		}
		left := i - (1 + p[i])
		right := i + (1 + p[i])
		for left >= 0 && right < len(p) && bytes[left] == bytes[right] {
			p[i]++
			left--
			right++
		}
		if i+p[i] > maxRight {
			maxRight = i + p[i]
			center = i
		}
		if p[i] > p[start] {
			start = i
		}
	}
	return s[(start-p[start])>>1 : (start+p[start])>>1]
}

/**
动态规划
时间复杂度：O(N^2)
空间复杂度：O(N)
*/
func longestPalindrome4(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	m := make([]bool, length)
	n := make([]bool, length)
	p := make([]bool, length)
	gap := 0
	var maxi, maxGap int
	for ; gap < length; gap++ {
		for i := gap; i < length; i++ {
			if s[i] == s[i-gap] {
				if gap < 3 {
					p[i-gap] = true
				} else {
					p[i-gap] = m[i-gap+1]
				}
				if p[i-gap] && maxGap < gap {
					maxi = i
					maxGap = gap
				}
			} else {
				p[i-gap] = false
			}
		}
		m, n, p = n, p, m
	}
	return s[maxi-maxGap : maxi+1]
}

/**
动态规划
时间复杂度：O(N^2)
空间复杂度：O((N*2)/2)
*/
func longestPalindrome3(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s)-i)
		dp[i][0] = true
	}
	gap := 1
	var maxi, maxGap int
	for ; gap < len(s); gap++ {
		for i := gap; i < len(s); i++ {
			if s[i] == s[i-gap] {
				if gap < 3 {
					dp[i-gap][gap] = true
				} else {
					dp[i-gap][gap] = dp[i-gap+1][gap-2]
				}
				if dp[i-gap][gap] && maxGap < gap {
					maxi = i
					maxGap = gap
				}
			}
		}
	}
	return s[maxi-maxGap : maxi+1]
}

/**
动态规划
时间复杂度：O(N^2)
空间复杂度：O(N^2)
*/
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	gap := 1
	var maxi, maxGap int
	for ; gap < len(s); gap++ {
		for i := gap; i < len(s); i++ {
			if s[i] == s[i-gap] {
				if gap < 3 {
					dp[i-gap][i] = true
				} else {
					dp[i-gap][i] = dp[i-gap+1][i-1]
				}
				if dp[i-gap][i] && maxGap < gap {
					maxi = i
					maxGap = gap
				}
			}
		}
	}
	return s[maxi-maxGap : maxi+1]
}

/**
中心扩散法
时间复杂度：O(N^2)
空间复杂度：O(1)
*/
func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}
	var max, idx int
	for i := 0; i < len(s); i++ {
		// 以i为中心辐射，回文子串为奇数
		left, right := i, i
		for left >= 0 && right < len(s) && (s)[left] == (s)[right] {
			left--
			right++
		}
		n1 := right - left - 1
		// 以i为中心辐射，回文子串为奇数
		left, right = i, i+1
		for left >= 0 && right < len(s) && (s)[left] == (s)[right] {
			left--
			right++
		}
		n2 := right - left - 1
		var n int
		if n1 > n2 {
			n = n1
		} else {
			n = n2
		}
		if n > max {
			idx = i
			max = n
		}
	}
	end := idx + 1 + max>>1
	return s[end-max : end]
}
