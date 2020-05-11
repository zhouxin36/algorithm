package 滑动窗口

import (
	"testing"
)

/**
LC3 https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/
func Test_leLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, test := range tests {
		if got := lengthOfLongestSubstring(test.input); got != test.want {
			t.Errorf("lengthOfLongestSubstring(%q) = %v", test.input, got)
		}
	}
}

/**
LC1425 https://leetcode.com/problems/constrained-subsequence-sum/
*/
func Test_constrainedSubsetSum(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, -2, -10, -5, 20}, 2, 23},
		{[]int{-8269, 3217, -4023, -4138, -683, 6455, -3621, 9242, 4015, -3790}, 1, 16091},
		{[]int{-1, -2, -3}, 1, -1},
		{[]int{10, 2, -10, 5, 20}, 2, 37},
	}
	for _, test := range tests {
		if got := constrainedSubsetSum(test.nums, test.k); got != test.want {
			t.Errorf("lengthOfLongestSubstring(%v,%v) = %v", test.nums, test.k, got)
		} else {
			//t.Logf("lengthOfLongestSubstring(%v,%v) = %v", test.nums, test.k, got)
		}
	}
}

/**
LC1438 https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
*/
func Test_longestSubarray(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{10, 1, 2, 4, 7, 2}, 5, 4},
		{[]int{8, 2, 4, 7}, 4, 2},
		{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0, 3},
	}
	for _, test := range tests {
		if got := longestSubarray(test.nums, test.k); got != test.want {
			t.Errorf("longestSubarray(%v,%v) = %v", test.nums, test.k, got)
		} else {
			//t.Logf("lengthOfLongestSubstring(%v,%v) = %v", test.nums, test.k, got)
		}
	}
}
