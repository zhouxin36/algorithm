package main

import "fmt"

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}) == 2)
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}) == 3)
}
func maxDistToClosest(seats []int) int {
	head, tag := -1, -1
	pref := 0
	max := 0
	for index, s := range seats {
		if s == 1 {
			if head == -1 {
				head = index
			}
			if index-pref > max {
				max = index - pref
			}
			pref = index
			tag = index
		}
	}
	tag = len(seats) - tag - 1
	max = max >> 1
	if max >= head && max >= tag {
		return max
	} else if head >= max && head >= tag {
		return head
	} else {
		return tag
	}
}
