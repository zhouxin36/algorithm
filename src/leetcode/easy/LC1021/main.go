package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeOuterParentheses("(()())(())") == "()()()")
	fmt.Println(removeOuterParentheses("(()())(())(()(()))") == "()()()()(())")
	fmt.Println(removeOuterParentheses("()()") == "")
}
func removeOuterParentheses(S string) string {
	var sb strings.Builder
	//var stack []byte
	var count int
	for _, s := range S {
		v := byte(s)
		if v == '(' {
			if count > 0 {
				sb.WriteByte(byte(s))
			}
			//stack = append(stack, v)
			count++
		} else {
			//stack = stack[:len(stack)-1]
			count--
			if count > 0 {
				sb.WriteByte(byte(s))
			}
		}
	}
	return sb.String()
}
