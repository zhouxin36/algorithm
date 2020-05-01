package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(5))
}
func firstBadVersion(n int) int {
	head, tag := 0, n
	for head <= tag {
		mid := (head + tag) >> 1
		if isBadVersion(mid) {
			tag = mid - 1
		} else {
			head = mid + 1
		}
	}
	return head
}
func isBadVersion(n int) bool {
	if n >= 4 {
		return true
	}
	return false
}
