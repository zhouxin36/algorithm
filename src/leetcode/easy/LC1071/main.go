package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX") == "TAUXX")
	fmt.Println(gcdOfStrings("ABCABC", "ABC") == "ABC")
	fmt.Println(gcdOfStrings("ABABAB", "ABAB") == "AB")
	fmt.Println(gcdOfStrings("LEET", "CODE") == "")
}
func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) > len(str2) {
		return doGcdOfStrings(str1, str2)
	} else {
		return doGcdOfStrings(str2, str1)
	}
}

func doGcdOfStrings(str1 string, str2 string) string {
	if str2 == "" {
		return str1
	} else {
		tmp := str1
		for strings.Index(tmp, str2) == 0 {
			tmp = tmp[len(str2):]
		}
		if tmp == str1 {
			return ""
		}
		return doGcdOfStrings(str2, tmp)
	}
}
