package main

import "fmt"

func main() {
	fmt.Println(findComplement(5) == 2)
	fmt.Println(findComplement(1) == 0)
}
func findComplement(num int) int {
	var pref int
	tmp := num
	for tmp != 0 {
		pref = tmp
		tmp &= tmp - 1
	}
	return pref<<1 - num - 1
}
