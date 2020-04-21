package main

import (
	bytes2 "bytes"
	"fmt"
)

func main() {
	//["e","l","l"]
	fmt.Println(commonChars([]string{"bella", "abel", "roller"}))
	//["c","o"]
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}
func commonChars(A []string) []string {
	var res []string
	if len(A) == 1 {
		for _, a := range A[0] {
			res = append(res, string(a))
		}
		return res
	}
	bytes := make([][]byte, len(A)-1)
	for index, a := range A[1:] {
		bytes[index] = []byte(a)
	}
	for _, v := range A[0] {
		flag := true
		for i := 0; i < len(bytes); i++ {
			indexByte := bytes2.IndexByte(bytes[i], byte(v))
			if indexByte == -1 {
				flag = false
				break
			}
			copy(bytes[i][indexByte:], bytes[i][indexByte+1:])
			bytes[i] = bytes[i][:len(bytes[i])-1]
		}
		if flag {
			res = append(res, string(v))
		}
	}
	return res
}
