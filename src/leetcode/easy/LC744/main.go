package main

import "fmt"

func main() {
	fmt.Println(nextGreatestLetter([]byte{'e', 'e', 'e', 'n', 'n'}, 'e') == 'n')
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a') == 'c')
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c') == 'f')
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd') == 'f')
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'g') == 'j')
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j') == 'c')
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'k') == 'c')
}
func nextGreatestLetter1(letters []byte, target byte) byte {
	for _, l := range letters {
		if l > target {
			return l
		}
	}
	return letters[0]
}
func nextGreatestLetter(letters []byte, target byte) byte {
	start := 0
	end := len(letters) - 1
	for start < end {
		mid := (start + end) >> 1
		if letters[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if letters[start] <= target {
		return letters[(start+1)%len(letters)]
	} else {
		return letters[start]
	}
}
