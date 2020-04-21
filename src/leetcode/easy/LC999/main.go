package main

import "fmt"

func main() {
	fmt.Println(numRookCaptures([][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', 'R', '.', '.', '.', 'p'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}}) == 3)
	fmt.Println(numRookCaptures([][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'B', 'R', 'B', 'p', '.', '.'}, {'.', 'p', 'p', 'B', 'p', 'p', '.', '.'}, {'.', 'p', 'p', 'p', 'p', 'p', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}}) == 0)
	fmt.Println(numRookCaptures([][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'p', 'p', '.', 'R', '.', 'p', 'B', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}}) == 3)
}
func numRookCaptures(board [][]byte) int {
	var l, r int
	var flag bool
	for i, v := range board {
		for j, jv := range v {
			if jv == 'R' {
				flag = true
				l = i
				r = j
				break
			}
		}
		if flag {
			break
		}
	}
	var count int
	for i := l + 1; i < len(board); i++ {
		if board[i][r] == '.' {
			continue
		}
		if board[i][r] == 'p' {
			count++
		}
		break
	}
	for i := l - 1; i >= 0; i-- {
		if board[i][r] == '.' {
			continue
		}
		if board[i][r] == 'p' {
			count++
		}
		break
	}
	for i := r + 1; i < len(board[0]); i++ {
		if board[l][i] == '.' {
			continue
		}
		if board[l][i] == 'p' {
			count++
		}
		break
	}
	for i := r - 1; i >= 0; i-- {
		if board[l][i] == '.' {
			continue
		}
		if board[l][i] == 'p' {
			count++
		}
		break
	}
	return count
}
