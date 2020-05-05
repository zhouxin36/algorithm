package main

import "fmt"

func main() {
	fmt.Println(tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}) == "A")
	fmt.Println(tictactoe([][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}) == "B")
	fmt.Println(tictactoe([][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}) == "Draw")
	fmt.Println(tictactoe([][]int{{0, 0}, {1, 1}}) == "Pending")
}
func tictactoe(moves [][]int) string {
	var ar, ac, br, bc [3]int
	var ad1, ad2, bd1, bd2 int
	for index, i := range moves {
		r, c := i[0], i[1]
		if index&1 == 0 {
			ar[r]++
			ac[c]++
			if r == c {
				ad1++
			}
			if r+c == 2 {
				ad2++
			}
			if ar[r] == 3 || ac[c] == 3 || ad1 == 3 || ad2 == 3 {
				return "A"
			}
		} else {
			br[r]++
			bc[c]++
			if r == c {
				bd1++
			}
			if r+c == 2 {
				bd2++
			}
			if br[r] == 3 || bc[c] == 3 || bd1 == 3 || bd2 == 3 {
				return "B"
			}
		}
	}
	if len(moves) == 9 {
		return "Draw"
	} else {
		return "Pending"
	}
}
