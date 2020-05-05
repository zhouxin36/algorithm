package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toHexspeak(257) == "IOI")
	fmt.Println(toHexspeak(3) == "ERROR")
}
func toHexspeak(num int) string {
	var sb strings.Builder
	for num != 0 {
		v := num & 0b1111
		switch v {
		case 0:
			sb.WriteByte('O')
		case 1:
			sb.WriteByte('I')
		case 10 - 15:
			sb.WriteByte(byte('A' + v - 10))
		default:
			return "ERROR"
		}
		num >>= 4
	}
	return sb.String()
}
