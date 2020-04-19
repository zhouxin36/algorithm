package main

import "fmt"

func main() {
	fmt.Println(minNumberOfFrogs("croak"))
}
func minNumberOfFrogs(croakOfFrogs string) int {
	c := 0
	r := 0
	o := 0
	a := 0
	k := 0
	inUse := 0
	answer := 0
	for _, d := range croakOfFrogs {
		switch d {
		case 'c':
			c++
			inUse++
		case 'r':
			r++
		case 'o':
			o++
		case 'a':
			a++
		case 'k':
			k++
			inUse--
		}
		if answer < inUse {
			answer = inUse
		}
		if (c < r) || (r < o) || (o < a) || (a < k) {
			return -1
		}
	}
	if inUse == 0 && c == r && c == o && c == a && c == k {
		return answer
	}
	return -1
}
