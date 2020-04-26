package main

import (
	"fmt"
	"strings"
)

func main() {
	//["girl","student"]
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	//["we","rock"]
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
}
func findOcurrences(text string, first string, second string) []string {
	t := strings.Split(text, " ")
	var res []string
	for i := 0; i < len(t); i++ {
		if t[i] == first && i+2 < len(t) && t[i+1] == second {
			res = append(res, t[i+2])
		}
	}
	return res
}
