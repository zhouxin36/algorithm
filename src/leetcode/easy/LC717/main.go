package main

import "fmt"

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(!isOneBitCharacter([]int{1, 1, 1, 0}))
}

/**
i为指针，当字符为0时移动一位，为1时移动两位
*/
func isOneBitCharacter1(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		i += bits[i] + 1
	}
	return i == len(bits)-1
}

/**
找到倒数第二个0，倒数第二个零后面只能全1或者为开始，计算剩下的1是否为偶数
*/
func isOneBitCharacter(bits []int) bool {
	i := len(bits) - 2
	for ; i >= 0 && bits[i] > 0; i-- {
	}
	return (len(bits)-i)%2 == 0
}
