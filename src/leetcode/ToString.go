package leetcode

import "encoding/json"

func IntToString(arr []int) string {
	bytes, _ := json.Marshal(arr)
	return string(bytes)
}
