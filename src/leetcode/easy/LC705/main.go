package main

import "fmt"

func main() {
	myHashSet := Constructor()
	myHashSet.Add(48)
	myHashSet.Add(41)
	myHashSet.Add(84)
	myHashSet.Add(82)
	myHashSet.Add(24)
	myHashSet.Add(7)
	myHashSet.Add(87)
	myHashSet.Add(81)
	myHashSet.Add(55)
	myHashSet.Add(19)
	myHashSet.Add(40)
	myHashSet.Add(68)
	myHashSet.Add(23)
	myHashSet.Add(23)
	myHashSet.Add(80)
	myHashSet.Add(53)
	myHashSet.Add(76)
	myHashSet.Add(95)
	myHashSet.Add(67)
	myHashSet.Add(31)
	myHashSet.Add(62)
	myHashSet.Add(73)
	myHashSet.Add(33)
	myHashSet.Add(28)
	myHashSet.Add(62)
	myHashSet.Add(81)
	myHashSet.Add(57)
	myHashSet.Add(11)
	myHashSet.Add(89)
	myHashSet.Add(28)
	myHashSet.Add(20)
	myHashSet.Add(77)
	myHashSet.Add(77)
	myHashSet.Add(52)
	myHashSet.Add(57)
	myHashSet.Add(88)
	myHashSet.Add(20)
	fmt.Println(myHashSet.Contains(48))
}

type MyHashSet struct {
	size      int
	container int
	table     []*Node
}
type Node struct {
	val  int
	next *Node
}

const InitContainer = 8
const ResetMultiple float32 = 1.75

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{0, InitContainer, make([]*Node, InitContainer)}
}

func (mhs *MyHashSet) nodeIndex(key int) int {
	return (mhs.container - 1) & key
}

func (mhs *MyHashSet) Add(key int) {
	index := mhs.nodeIndex(key)
	v := mhs.table[index]
	if v == nil {
		mhs.table[index] = &Node{val: key}
		mhs.size++
	} else if v.val == key {
		return
	} else {
		for v.next != nil {
			if v.next.val == key {
				return
			}
			v = v.next
		}
		v.next = &Node{val: key}
		mhs.size++
	}
	if float32(mhs.size) >= float32(mhs.container)*ResetMultiple {
		mhs.reset()
	}
}

func (mhs *MyHashSet) Remove(key int) {
	index := mhs.nodeIndex(key)
	v := mhs.table[index]
	if v == nil {
		return
	}
	if v.val == key {
		mhs.table[index] = v.next
		mhs.size--
		return
	}
	for v != nil && v.next != nil {
		if v.next.val == key {
			v.next = v.next.next
			mhs.size--
			return
		}
		v = v.next
	}
}
func (mhs *MyHashSet) reset() {
	old := mhs.container
	mhs.container = old << 1
	newTable := make([]*Node, mhs.container)
	for i := 0; i < len(mhs.table); i++ {
		if mhs.table[i] == nil {
			continue
		}
		var loHead, loTail, hiHead, hiTail, next *Node
		next = mhs.table[i]
		for next != nil {
			if mhs.nodeIndex(next.val) == i {
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
	mhs.table = newTable
}

/** Returns true if this set contains the specified element */
func (mhs *MyHashSet) Contains(key int) bool {
	index := mhs.nodeIndex(key)
	v := mhs.table[index]
	for v != nil {
		if v.val == key {
			return true
		}
		v = v.next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
