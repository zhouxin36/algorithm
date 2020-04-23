package main

import "fmt"

func main() {
	fmt.Println(divisorGame(2))
	fmt.Println(!divisorGame(3))
}
func divisorGame(N int) bool {
	return N&1 == 0
}
