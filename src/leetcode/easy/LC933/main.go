package main

import "fmt"

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
}

type RecentCounter struct {
	Head []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}
func (this *RecentCounter) Ping(t int) int {
	num := t - 3000
	this.Head = append(this.Head, t)
	if num > 0 {
		for true {
			if len(this.Head) > 0 && this.Head[0] < num {
				this.Head = this.Head[1:]
			} else {
				break
			}
		}
	}
	return len(this.Head)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
