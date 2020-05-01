package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko") == 1)
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon") == 2)
	fmt.Println(maxNumberOfBalloons("leetcode") == 0)
}
func maxNumberOfBalloons(text string) int {
	//balloon
	balloon := map[byte]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}
	for _, v := range text {
		if _, ok := balloon[byte(v)]; ok {
			balloon[byte(v)]++
		}
	}
	res := math.MaxInt32
	for k, v := range balloon {
		switch k {
		case 'a', 'b', 'n':
			if res > v {
				res = v
			}
		case 'l', 'o':
			if res > v>>1 {
				res = v >> 1
			}
		}
	}
	return res
}
