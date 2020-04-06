package main

import "fmt"

func main() {
	myHashMap := Constructor()
	myHashMap.Put(4, 30)
	myHashMap.Put(41, 24)
	myHashMap.Put(82, 59)
	myHashMap.Put(92, 22)
	myHashMap.Put(34, 41)
	myHashMap.Put(59, 59)
	myHashMap.Put(20, 35)
	myHashMap.Put(24, 74)
	myHashMap.Put(95, 10)
	myHashMap.Put(95, 10)
	myHashMap.Remove(59)
	myHashMap.Put(66, 1)
	myHashMap.Put(20, 46)
	myHashMap.Remove(59)
	myHashMap.Remove(2)
	myHashMap.Put(28, 13)
	myHashMap.Put(45, 16)
	myHashMap.Put(12, 8)
	myHashMap.Remove(24)
	myHashMap.Put(27, 61)
	myHashMap.Put(43, 57)
	myHashMap.Put(72, 24)
	myHashMap.Put(72, 24)
	myHashMap.Remove(20)
	myHashMap.Put(95, 75)
	myHashMap.Remove(20)
	myHashMap.Put(64, 8)
	myHashMap.Put(36, 61)
	myHashMap.Put(95, 53)
	myHashMap.Put(76, 34)
	myHashMap.Remove(24)
	myHashMap.Put(55, 68)
	myHashMap.Remove(21)
	myHashMap.Put(75, 71)
	myHashMap.Put(43, 36)
	myHashMap.Put(50, 3)
	myHashMap.Put(4, 97)
	myHashMap.Put(44, 57)
	myHashMap.Put(20, 1)
	myHashMap.Put(4, 66)
	myHashMap.Put(8, 71)
	myHashMap.Put(34, 41)
	myHashMap.Put(75, 64)
	myHashMap.Put(98, 47)
	myHashMap.Put(12, 45)
	myHashMap.Put(93, 69)
	myHashMap.Put(93, 10)
	myHashMap.Put(64, 66)
	myHashMap.Put(72, 99)
	myHashMap.Put(50, 97)
	myHashMap.Put(48, 65)
	myHashMap.Remove(80)
	myHashMap.Put(46, 13)
	myHashMap.Put(60, 15)
	myHashMap.Put(42, 3)
	myHashMap.Put(29, 18)
	myHashMap.Put(95, 70)
	myHashMap.Put(0, 80)
	myHashMap.Put(6, 15)
	myHashMap.Remove(73)
	myHashMap.Put(26, 45)
	myHashMap.Put(60, 91)
	myHashMap.Put(13, 32)
	fmt.Println(myHashMap.Get(75))
}

type MyHashMap struct {
	size      int
	container int
	table     []*Node
}
type Node struct {
	key  int
	val  int
	next *Node
}

const InitContainer = 8
const ResetMultiple float32 = 1.75

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{0, InitContainer, make([]*Node, InitContainer)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	index := this.nodeIndex(key)
	v := this.table[index]
	if v == nil {
		this.table[index] = &Node{key: key, val: value}
		this.size++
	} else if v.key == key {
		v.val = value
		return
	} else {
		for v.next != nil {
			if v.next.key == key {
				v.next.val = value
				return
			}
			v = v.next
		}
		v.next = &Node{key: key, val: value}
		this.size++
	}
	if float32(this.size) >= float32(this.container)*ResetMultiple {
		this.reset()
	}
}

func (this *MyHashMap) nodeIndex(key int) int {
	return (this.container - 1) & key
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	index := this.nodeIndex(key)
	v := this.table[index]
	for v != nil {
		if v.key == key {
			return v.val
		}
		v = v.next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	index := this.nodeIndex(key)
	v := this.table[index]
	if v == nil {
		return
	}
	if v.key == key {
		this.table[index] = v.next
		this.size--
		return
	}
	for v != nil && v.next != nil {
		if v.next.key == key {
			v.next = v.next.next
			this.size--
			return
		}
		v = v.next
	}
}

func (this *MyHashMap) reset() {
	old := this.container
	this.container = old << 1
	newTable := make([]*Node, this.container)
	for i := 0; i < len(this.table); i++ {
		if this.table[i] == nil {
			continue
		}
		var loHead, loTail, hiHead, hiTail, next *Node
		next = this.table[i]
		for next != nil {
			if this.nodeIndex(next.key) == i {
				if loHead == nil {
					loHead = next
				} else {
					loTail.next = next
				}
				loTail = next
			} else {
				if hiHead == nil {
					hiHead = next
				} else {
					hiTail.next = next
				}
				hiTail = next
			}
			next = next.next
		}
		if loTail != nil {
			loTail.next = nil
			newTable[i] = loHead
		}
		if hiTail != nil {
			hiTail.next = nil
			newTable[old+i] = hiHead
		}
	}
	this.table = newTable
}
