package main

import "fmt"

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
	fmt.Println(!isLongPressedName("pyplrz", "ppyypllr"))
	fmt.Println(!isLongPressedName("saeed", "ssaaedd"))
	fmt.Println(isLongPressedName("leelee", "lleeelee"))
	fmt.Println(isLongPressedName("laiden", "laiden"))
}
func isLongPressedName(name string, typed string) bool {
	if name[0] != typed[0] {
		return false
	}
	nIndex, tIndex := 1, 1
	for ; tIndex < len(typed); tIndex++ {
		if nIndex < len(name) && name[nIndex] == typed[tIndex] {
			nIndex++
		} else if name[nIndex-1] == typed[tIndex] {

		} else {
			return false
		}
	}
	return nIndex == len(name)
}
