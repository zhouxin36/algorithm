package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("babadada") == "adada")
	fmt.Println(longestPalindrome("abcdbbfcba") == "bb")
	fmt.Println(longestPalindrome("bb") == "bb")
	fmt.Println(longestPalindrome("abcda") == "a")
	fmt.Println(longestPalindrome("abcba") == "abcba")
	fmt.Println(longestPalindrome("babad") == "bab")
	fmt.Println(longestPalindrome("cbbd") == "bb")
}

/**
Manacher 算法
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
			if maxRight-i > p[mirror] {
				p[i] = p[mirror]
			} else {
				p[i] = maxRight - i
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

	var start, maxGap int
	for gap := 0; gap < length; gap++ {
		for i := 0; i < length-gap; i++ {
			if s[i+gap] == s[i] {
				if gap < 3 {
					// gap为0时，子串为1；为1时，子串为2
					p[i] = true
				} else {
					// i+1为i缩短一个字符，m与p差两个gap；表示除去头和尾的子串
					p[i] = m[i+1]
				}
				if p[i] && maxGap < gap {
					start = i
					maxGap = gap
				}
			} else {
				p[i] = false
			}
		}
		m, n, p = n, p, m
	}
	return s[start : start+maxGap+1]
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
	var r, l, left, right int
	for left < len(s) {
		// 以i为中心扩散，匹配相同的字符
		for right+1 < len(s) && (s)[left] == (s)[right+1] {
			right++
		}
		// 以left，right，向两边扩散
		for left > 0 && right < len(s)-1 && (s)[left-1] == (s)[right+1] {
			left--
			right++
		}
		if r-l < right-left {
			r, l = right, left
		}
		// 取中间值
		left = (left+right)>>1 + 1
		right = left
	}
	return s[l : r+1]
}
