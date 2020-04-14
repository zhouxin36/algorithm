package main

import (
	"fmt"
	"strings"
)

func main() {
	//["sweet","sour"]
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	//banana
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}
func uncommonFromSentences(A string, B string) []string {
	m := map[string]int{}
	sA := strings.Split(A, " ")
	sB := strings.Split(B, " ")
	for _, s := range sA {
		m[s]++
	}
	for _, s := range sB {
		m[s]++
	}
	var result []string
	for k, v := range m {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}
