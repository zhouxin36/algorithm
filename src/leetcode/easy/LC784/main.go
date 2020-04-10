package main

import "fmt"

func main() {
	//["a1b2", "a1B2", "A1b2", "A1B2"]
	fmt.Println(letterCasePermutation("a1b2"))
	//["3z4", "3Z4"]
	fmt.Println(letterCasePermutation("3z4"))
	//["12345"]
	fmt.Println(letterCasePermutation("12345"))
}
func letterCasePermutation(S string) []string {
	return letterCasePermutation1(S)
}
func letterCasePermutation1(S string) []string {
	var indexes []int
	for index, s := range S {
		if s >= '0' && s <= '9' {

		} else {
			indexes = append(indexes, index)
		}
	}
	length := len(indexes)
	if length == 0 {
		return []string{S}
	}
	return getString(length, indexes, S)
}

func getString(length int, indexes []int, s string) []string {
	num := 1 << uint8(length)
	var result []string
	for i := 0; i < num; i++ {
		bytes := []byte(s)
		n := i
		for _, index := range indexes {
			if n&1 != 0 {
				if bytes[index] >= 'a' && bytes[index] <= 'z' {
					bytes[index] -= 32
				} else {
					bytes[index] += 32
				}
			}
			n >>= 1
		}
		result = append(result, string(bytes))
	}
	return result
}
