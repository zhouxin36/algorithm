package main

import "fmt"

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd") == "dc-ba")
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj") == "j-Ih-gfE-dCba")
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!") == "Qedo1ct-eeLg=ntse-T!")
}
func reverseOnlyLetters(S string) string {
	head, tag := 0, len(S)-1
	A := []byte(S)
	for head < tag {
		for head <= tag && !isA(A[head]) {
			head++
		}
		for tag >= head && !isA(A[tag]) {
			tag--
		}
		if head >= tag {
			break
		}
		A[head] ^= A[tag]
		A[tag] ^= A[head]
		A[head] ^= A[tag]
		head++
		tag--
	}
	return string(A)
}
func isA(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z')
}
