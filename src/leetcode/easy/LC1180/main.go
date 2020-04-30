package main

import "fmt"

func main() {
	fmt.Println(countSubstring("aaaba") == 8)
	fmt.Println(countSubstring("aaaaaaaaaa") == 55)
}
func countSubstring(str string) int {
	var res int
	count := 1
	pref := str[0]
	for i := 1; i < len(str); i++ {
		if str[i] == pref {
			count++
		} else {
			v := 0
			for j := 1; j <= count; j++ {
				v += j
			}
			res += v
			count = 1
			pref = str[i]
		}
	}
	v := 0
	for j := 1; j <= count; j++ {
		v += j
	}
	res += v
	return res
}
