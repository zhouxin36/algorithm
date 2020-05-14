package main

import "fmt"

func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("app"))     // returns false
	fmt.Println(trie.Search("apple"))   // returns true
	fmt.Println(trie.StartsWith("app")) // returns true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // returns true
}

const SIZE = 26

type Trie struct {
	m    [26]*Trie
	flag bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{m: [SIZE]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	head := this
	for i := 0; i < len(word); i++ {
		v := head.m[word[i]-'a']
		if v != nil {
			head = v
		} else {
			head.m[word[i]-'a'] = &Trie{m: [SIZE]*Trie{}}
			head = head.m[word[i]-'a']
		}
		if i == len(word)-1 {
			head.flag = true
		}
	}

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	head := this
	for i := 0; i < len(word); i++ {
		v := head.m[word[i]-'a']
		if v != nil {
			head = v
		} else {
			return false
		}
	}
	return head.flag
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	head := this.m
	for i := 0; i < len(prefix); i++ {
		v := head[prefix[i]-'a']
		if v != nil {
			head = v.m
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
