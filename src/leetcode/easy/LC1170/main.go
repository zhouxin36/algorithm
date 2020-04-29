package main

import (
	"fmt"
	"sort"
)

func main() {
	//[6,1,1,2,3,3,3,1,3,2]
	fmt.Println(numSmallerByFrequency([]string{"bba", "abaaaaaa", "aaaaaa", "bbabbabaab", "aba", "aa", "baab", "bbbbbb", "aab", "bbabbaabb"}, []string{"aaabbb", "aab", "babbab", "babbbb", "b", "bbbbbbbbab", "a", "bbbbbbbbbb", "baaabbaab", "aa"}))
	//[1]
	fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	//[1,2]
	fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}
func numSmallerByFrequency(queries []string, words []string) []int {
	qA := make([]int, len(queries))
	wA := make([]int, len(words))
	for index, v := range queries {
		curByte := 'z'
		var count int
		for _, va := range v {
			if va == curByte {
				count++
			} else if curByte > va {
				curByte = va
				count = 1
			}
		}
		qA[index] = count
	}
	for index, v := range words {
		curByte := 'z'
		var count int
		for _, va := range v {
			if va == curByte {
				count++
			} else if curByte > va {
				curByte = va
				count = 1
			}
		}
		wA[index] = count
	}
	fmt.Println(qA)
	sort.Sort(sort.Reverse(sort.IntSlice(wA)))
	fmt.Println(wA)
	res := make([]int, len(queries))
	for i, v := range qA {
		head, tag := 0, len(wA)-1
		for head < tag {
			mid := (head + tag) >> 1
			if wA[mid] > v {
				head = mid + 1
			} else if wA[mid] < v {
				tag = mid - 1
			} else {
				head = mid
				break
			}
		}
		for ; head >= 0; head-- {
			if wA[head] > v {
				break
			}
		}
		for ; head >= 0 && head < len(wA); head++ {
			if wA[head] <= v {
				break
			}
		}
		if head <= 0 {
			res[i] = 0
		} else if head >= len(wA) {
			res[i] = len(wA)
		} else {
			res[i] = head
		}
	}
	return res
}
