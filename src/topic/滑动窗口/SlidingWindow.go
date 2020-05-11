package 滑动窗口

/**
LC3 https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var arr [128]bool
	var count int
	for i, j := 0, 0; i < len(s) && j < len(s); {
		if arr[s[j]] {
			arr[s[i]] = false
			i++
		} else {
			arr[s[j]] = true
			if j-i > count {
				count = j - i
			}
			j++
		}
	}
	return count + 1
}

/**
LC1425 https://leetcode.com/problems/constrained-subsequence-sum/
*/
func constrainedSubsetSum(nums []int, k int) int {
	res := nums[0]
	var q []int
	for i, _ := range nums {
		if len(q) != 0 {
			nums[i] += q[0]
		}
		if res < nums[i] {
			res = nums[i]
		}
		for len(q) > 0 && nums[i] > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		if nums[i] > 0 {
			q = append(q, nums[i])
		}
		if i >= k && len(q) > 0 && q[0] == nums[i-k] {
			q = q[1:]
		}
	}
	return res
}

/**
LC1438 https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
*/
func longestSubarray(nums []int, limit int) int {
	var max, min []int
	var i, j int
	for ; i < len(nums); i++ {
		for len(max) > 0 && nums[i] > max[len(max)-1] {
			max = max[:len(max)-1]
		}
		for len(min) > 0 && nums[i] < min[len(min)-1] {
			min = min[:len(min)-1]
		}
		max = append(max, nums[i])
		min = append(min, nums[i])
		if max[0]-min[0] > limit {
			if max[0] == nums[j] {
				max = max[1:]
			}
			if min[0] == nums[j] {
				min = min[1:]
			}
			j++
		}
	}
	return i - j
}
