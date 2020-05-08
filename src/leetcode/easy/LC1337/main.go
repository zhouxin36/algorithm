package main

import (
	"fmt"
	"sort"
)

func main() {
	//[4,6,0,1,2,3]
	fmt.Println(kWeakestRows([][]int{{1, 1, 0}, {1, 1, 0}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}, {1, 1, 1}, {1, 0, 0}}, 6))
	//2,0,3
	fmt.Println(kWeakestRows([][]int{{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}, 3))
	//0,2
	fmt.Println(kWeakestRows([][]int{{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0}}, 2))
}

type help struct {
	index int
	val   int
}

func kWeakestRows(mat [][]int, k int) []int {
	heap := make([]help, k)
	var length int
	for i, v := range mat {
		var n int
		for _, v2 := range v {
			if v2 == 0 {
				break
			} else {
				n++
			}
		}
		if length < k {
			heap[length] = help{index: i, val: n}
			length++
			// balance
			if length == k {
				for i := (k >> 1) - 1; i >= 0; i-- {
					balance(heap, i)
				}
			}
		} else if heap[0].val > n {
			heap[0] = help{index: i, val: n}
			balance(heap, 0)
		}
	}
	sort.Slice(heap, func(i, j int) bool {
		return heap[i].index < heap[j].index
	})
	sort.SliceStable(heap, func(i, j int) bool {
		return heap[i].val < heap[j].val
	})
	res := make([]int, k)
	for i, v := range heap {
		res[i] = v.index
	}
	return res
}

func balance(heap []help, i int) {
	for 2*i < len(heap)-1 {
		index := 2*i + 1
		if index+1 < len(heap) &&
			((heap[index].val < heap[index+1].val) ||
				(heap[index].val == heap[index+1].val && heap[index].index < heap[index+1].index)) {
			index++
		}
		if heap[i].val > heap[index].val ||
			(heap[i].val == heap[index].val && heap[i].index > heap[index].index) {
			break
		}
		heap[i], heap[index] = heap[index], heap[i]
		i = index
	}
}
