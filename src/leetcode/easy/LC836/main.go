package main

import (
	"fmt"
)

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	fmt.Println(!isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
}
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 2 在1左边,右边,上边，下边
	return !(rec2[2] <= rec1[0] || rec2[0] >= rec1[2] || rec2[1] >= rec1[3] || rec2[3] <= rec1[1])
}

/**
重叠区域大于零
*/
//func isRectangleOverlap(rec1 []int, rec2 []int) bool {
//return (math.min(rec1[2], rec2[2]) > Math.max(rec1[0], rec2[0]) && // width > 0
//Math.min(rec1[3], rec2[3]) > Math.max(rec1[1], rec2[1]));  // height > 0
//}
