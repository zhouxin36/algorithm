package main

import (
	"fmt"
)

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(!buddyStrings("ab", "ab"))
	fmt.Println(buddyStrings("aa", "aa"))
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
	fmt.Println(!buddyStrings("", "aa"))
}
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	m := map[byte]struct{}{}
	flag := false
	pref, after := -1, -1
	count := 0
	for i := 0; i < len(A); i++ {
		if !flag {
			if _, ok := m[A[i]]; ok {
				flag = true
			} else {
				m[A[i]] = struct{}{}
			}
		}
		if A[i] == B[i] {
			continue
		}
		if count == 0 {
			pref = i
			count++
		} else if count == 1 {
			after = i
			count++
		} else {
			return false
		}
	}
	if count == 0 {
		return flag
	}
	if count == 1 {
		return false
	}
	if A[pref] == B[after] && B[pref] == A[after] {
		return true
	}
	return false
}
