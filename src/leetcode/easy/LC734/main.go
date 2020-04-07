package main

import "fmt"

func main() {
	fmt.Println(areSentencesSimilar([]string{"great", "acting", "skills"},
		[]string{"fine", "drama", "talent"}, [][]string{{"great", "fine"},
			{"acting", "drama"}, {"skills", "talent"}}))

}

/**
暴力解
*/
func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
	return false
}
