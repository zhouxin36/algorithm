package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	//["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
	fmt.Println(reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}))
}
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(x int, y int) bool {
		x1 := strings.SplitN(logs[x], " ", 2)
		y1 := strings.SplitN(logs[y], " ", 2)
		dx := unicode.IsDigit(rune(x1[1][0]))
		dy := unicode.IsDigit(rune(y1[1][0]))
		if !dx && !dy {
			return strings.Compare(x1[1], y1[1]) != 1
		}
		if dx {
			return false
		} else {
			return true
		}
	})
	return logs
}
