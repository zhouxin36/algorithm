package main

import "fmt"

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}
	var kthLargest = Constructor(k, arr)
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(5))
	fmt.Println(kthLargest.Add(10))
	fmt.Println(kthLargest.Add(9))
	fmt.Println(kthLargest.Add(4))

	k = 7
	arr = []int{-10, 1, 3, 1, 4, 10, 3, 9, 4, 5, 1}
	kthLargest = Constructor(k, arr)
	fmt.Println(kthLargest.Add(3))
}

type KthLargest struct {
	size int
	k    int
	arr  []int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{0, k, []int{}}
	i := 0
	for ; i < len(nums); i++ {
		kl.Add(nums[i])
	}

	return kl
}

func balance(kl *KthLargest) {
	n := 0
	for true {
		index := 2*n + 1
		if index >= kl.size {
			return
		}
		if index < kl.size-1 && kl.arr[index] > kl.arr[index+1] {
			index++
		}
		if kl.arr[n] <= kl.arr[index] {
			return
		}
		kl.arr[n] ^= kl.arr[index]
		kl.arr[index] ^= kl.arr[n]
		kl.arr[n] ^= kl.arr[index]
		n = index
	}
}

func initMethod(kl *KthLargest) {
	n := kl.size - 1
	for n > 0 {
		p := (n - 1) / 2
		if kl.arr[n] >= kl.arr[p] {
			return
		} else {
			kl.arr[n] ^= kl.arr[p]
			kl.arr[p] ^= kl.arr[n]
			kl.arr[n] ^= kl.arr[p]
			n = p
		}
	}
}

func (kl *KthLargest) Add(val int) int {
	if kl.size < kl.k {
		kl.arr = append(kl.arr, val)
		//kl.arr[kl.size] = val
		kl.size++
		initMethod(kl)
	} else {
		if val > kl.arr[0] {
			kl.arr[0] = val
			balance(kl)
		}
	}
	return kl.arr[0]
}
