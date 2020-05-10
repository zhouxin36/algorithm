package main

import "strings"

func main() {
}
func generateTheString(n int) string {
	var sb strings.Builder
	if n&1 == 0 {
		sb.WriteByte('a')
		n--
	}
	for i := 0; i < n; i++ {
		sb.WriteByte('b')
	}
	return sb.String()
}
