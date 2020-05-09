package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	//0,1,2,4,8,3,5,6,7
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	//1,2,4,8,16,32,64,128,256,512,1024
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
	//10000,10000
	fmt.Println(sortByBits([]int{10000, 10000}))
	//2,3,5,17,7,11,13,19
	fmt.Println(sortByBits([]int{2, 3, 5, 7, 11, 13, 17, 19}))
	//10,100,10000,1000
	fmt.Println(sortByBits([]int{10, 100, 1000, 10000}))
}
func sortByBits(arr []int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		vi, vj := bits.OnesCount16(uint16(arr[i])), bits.OnesCount16(uint16(arr[j]))
		if vi > vj {
			return false
		} else if vi < vj {
			return true
		} else {
			return arr[i] < arr[j]
		}
	})
	return arr
}

// 汉明重量算法
const m1 = 0x5555555555555555  //binary: 0101...
const m2 = 0x3333333333333333  //binary: 00110011..
const m4 = 0x0f0f0f0f0f0f0f0f  //binary:  4 zeros,  4 ones ...
const m8 = 0x00ff00ff00ff00ff  //binary:  8 zeros,  8 ones ...
const m16 = 0x0000ffff0000ffff //binary: 16 zeros, 16 ones ...
const m32 = 0x00000000ffffffff //binary: 32 zeros, 32 ones ...
const h01 = 0x0101010101010101 //the sum of 256 to the power of 0,1,2,3...

//This is a naive implementation, shown for comparison,
//and to help in understanding the better functions.
//It uses 24 arithmetic operations (shift, add, and).
func PopCount1(x int) int {
	x = (x & m1) + ((x >> 1) & m1)    //put count of each  2 bits into those  2 bits
	x = (x & m2) + ((x >> 2) & m2)    //put count of each  4 bits into those  4 bits
	x = (x & m4) + ((x >> 4) & m4)    //put count of each  8 bits into those  8 bits
	x = (x & m8) + ((x >> 8) & m8)    //put count of each 16 bits into those 16 bits
	x = (x & m16) + ((x >> 16) & m16) //put count of each 32 bits into those 32 bits
	x = (x & m32) + ((x >> 32) & m32) //put count of each 64 bits into those 64 bits
	return x
}

//This uses fewer arithmetic operations than any other known
//implementation on machines with slow multiplication.
//It uses 17 arithmetic operations.
func PopCount2(x int) int {
	x -= (x >> 1) & m1             //put count of each 2 bits into those 2 bits
	x = (x & m2) + ((x >> 2) & m2) //put count of each 4 bits into those 4 bits
	x = (x + (x >> 4)) & m4        //put count of each 8 bits into those 8 bits
	x += x >> 8                    //put count of each 16 bits into their lowest 8 bits
	x += x >> 16                   //put count of each 32 bits into their lowest 8 bits
	x += x >> 32                   //put count of each 64 bits into their lowest 8 bits
	return x & 0xff
}

//This uses fewer arithmetic operations than any other known
//implementation on machines with fast multiplication.
//It uses 12 arithmetic operations, one of which is a multiply.
func PopCount3(x int) int {
	x -= (x >> 1) & m1             //put count of each 2 bits into those 2 bits
	x = (x & m2) + ((x >> 2) & m2) //put count of each 4 bits into those 4 bits
	x = (x + (x >> 4)) & m4        //put count of each 8 bits into those 8 bits
	return (x * h01) >> 56         //returns left 8 bits of x + (x<<8) + (x<<16) + (x<<24) + ...
}
