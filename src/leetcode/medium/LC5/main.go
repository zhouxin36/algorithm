package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("abcba") == "abcba")
	fmt.Println(longestPalindrome("abcda") == "a")
	fmt.Println(longestPalindrome("babad") == "bab")
	fmt.Println(longestPalindrome("cbbd") == "bb")
}

/**
动态规划
时间复杂度：O(N^2)
空间复杂度：O(N*lgN)
*/
func longestPalindrome(s string) string {
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
