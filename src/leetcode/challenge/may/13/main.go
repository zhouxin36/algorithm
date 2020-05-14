package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeKdigits("100", 1) == "0")
	fmt.Println(removeKdigits("1234567890", 9) == "0")
	fmt.Println(removeKdigits("10", 1) == "0")
	fmt.Println(removeKdigits("5337", 2) == "33")
	fmt.Println(removeKdigits("4321", 2) == "21")
	fmt.Println(removeKdigits("10200", 1) == "200")
	fmt.Println(removeKdigits("1111111", 3) == "1111")
	fmt.Println(removeKdigits("0", 0) == "0")
	fmt.Println(removeKdigits("1432239", 3) == "1223")
	fmt.Println(removeKdigits("212", 1) == "12")
	fmt.Println(removeKdigits("1432219", 3) == "1219")
	fmt.Println(removeKdigits("10", 2) == "0")
}
func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	} else if k == 0 {
		return num
	}
	var sb []byte
	i := 0
	for i < len(num) && k > 0 {
		if len(sb) == 0 {
			sb = append(sb, num[i])
			i++
			continue
		}
		if sb[len(sb)-1] > num[i] {
			sb = sb[:len(sb)-1]
			k--
		} else {
			sb = append(sb, num[i])
			i++
		}
	}
	sb = append(sb, num[i:]...)
	i = 0
	for ; i < len(sb)-1; i++ {
		if sb[i] != '0' {
			break
		}
	}
	//fmt.Println(string(sb[i:len(sb)-k]))
	return string(sb[i : len(sb)-k])
}
