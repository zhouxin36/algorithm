package main

import "fmt"

func main() {
	//[0,4,1,3,2]
	fmt.Println(diStringMatch("IDID"))
	//[0,1,2,3]
	fmt.Println(diStringMatch("III"))
	//[3,2,0,1]
	fmt.Println(diStringMatch("DDI"))
}
func diStringMatch(S string) []int {
	head, tag := 0, len(S)
	var result []int
	for _, s := range S {
		if s == 'I' {
			result = append(result, head)
			head++
		} else {
			result = append(result, tag)
			tag--
		}
	}
	return append(result, head)
}
