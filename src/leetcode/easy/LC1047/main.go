package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("abba"))
}
func removeDuplicates(S string) string {
	bs := []byte(S)
	for true {
		var flag bool
		for i := 1; i < len(bs); i++ {
			if bs[i] == bs[i-1] {
				flag = true
				copy(bs[i-1:], bs[i+1:])
				bs = bs[:len(bs)-2]
			}
		}
		if !flag {
			break
		}
	}
	return string(bs)
}
