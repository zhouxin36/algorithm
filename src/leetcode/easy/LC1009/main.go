package main

import "fmt"

func main() {
	fmt.Println(bitwiseComplement(5) == 2)
	fmt.Println(bitwiseComplement(7) == 0)
	fmt.Println(bitwiseComplement(10) == 5)
}
func bitwiseComplement(N int) int {
	b := 1
	for b < N {
		b = (b << 1) + 1
	}
	return N ^ b
}
