package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3) == "III")
	fmt.Println(intToRoman(4) == "IV")
	fmt.Println(intToRoman(9) == "IX")
	fmt.Println(intToRoman(58) == "LVIII")
	fmt.Println(intToRoman(1994) == "MCMXCIV")
}
func intToRoman(num int) string {
	var res strings.Builder
	for num != 0 {
		switch {
		case num < 4:
			res.WriteString("I")
			num--
		case num < 9:
			if num >= 5 {
				res.WriteString("V")
				num -= 5
			} else {
				res.WriteString("IV")
				num -= 4
			}
		case num < 40:
			if num >= 10 {
				res.WriteString("X")
				num -= 10
			} else {
				res.WriteString("IX")
				num -= 9
			}
		case num < 90:
			if num >= 50 {
				res.WriteString("L")
				num -= 50
			} else {
				res.WriteString("XL")
				num -= 40
			}
		case num < 400:
			if num >= 100 {
				res.WriteString("C")
				num -= 100
			} else {
				res.WriteString("XC")
				num -= 90
			}
		case num < 900:
			if num >= 500 {
				res.WriteString("D")
				num -= 500
			} else {
				res.WriteString("CD")
				num -= 400
			}
		default:
			if num >= 1000 {
				res.WriteString("M")
				num -= 1000
			} else {
				res.WriteString("CM")
				num -= 900
			}
		}
	}
	fmt.Println(res.String())
	return res.String()
}
