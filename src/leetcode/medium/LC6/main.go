package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 1) == "PAYPALISHIRING")
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
}
func convert(s string, numRows int) string {
	var sb strings.Builder
	var gap int
	if numRows == 1 {
		gap = 1
	} else {
		gap = 2 * (numRows - 1)
	}

	for row := 0; row < numRows; row++ {
		for i := row; i < len(s); {
			sb.WriteByte(s[i])
			if row != 0 && row != numRows-1 {
				n := i + 2*(numRows-1-row)
				if n < len(s) {
					sb.WriteByte(s[n])
				}
			}
			i += gap
		}
	}
	return sb.String()
}
