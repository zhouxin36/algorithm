package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("+1") == 1)
	fmt.Println(myAtoi(" ") == 0)
	fmt.Println(myAtoi("3.14159") == 3)
	fmt.Println(myAtoi("42") == 42)
	fmt.Println(myAtoi("   -42") == -42)
	fmt.Println(myAtoi("4193 with words") == 4193)
	fmt.Println(myAtoi("words and 987") == 0)
	fmt.Println(myAtoi("-91283472332") == -2147483648)
}
func myAtoi(str string) int {
	var i int
for1:
	for ; i < len(str); i++ {
		switch str[i] {
		case ' ':
			continue
		case '-', '+', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			break for1
		default:
			return 0
		}
	}
	var flag bool
	if i >= len(str) {
		return 0
	} else if str[i] == '-' {
		flag = true
		i++
	} else if str[i] == '+' {
		i++
	}
	var res int64
for2:
	for ; i < len(str); i++ {
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			res = 10*res + int64(str[i]-'0')
			if res > math.MaxInt32 {
				break for2
			}
		default:
			break for2
		}
	}
	if flag {
		res *= -1
		if res < math.MinInt32 {
			return math.MinInt32
		} else {
			return int(res)
		}
	} else {
		if res > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return int(res)
		}
	}
}
