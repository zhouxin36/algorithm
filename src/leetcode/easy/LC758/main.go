package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(boldWords([]string{"ab", "bc"}, "aabcd") == "a<b>abc</b>d")
}
func boldWords(words []string, s string) string {
	flag := make([]bool, len(s))
	for _, word := range words {
		for i := range s {
			f := true
			for j, ws := range word {
				if uint8(ws) != s[j+i] {
					f = false
					break
				}
			}
			if f {
				for a := i; a < i+len(word); a++ {
					flag[a] = true
				}
			}
		}
	}
	var result strings.Builder
	if flag[0] {
		result.WriteString("<br>")
	}
	length := len(flag)
	for i := 0; i < length; i++ {
		result.WriteByte(s[i])
		if !flag[i] && i < length-1 && flag[i+1] {
			result.WriteString("<br>")
		}
		if flag[i] && i < length-1 && !flag[i+1] {
			result.WriteString("</br>")
		}
	}
	fmt.Println(result.String())
	return result.String()
}
