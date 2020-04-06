package main

import "fmt"

func main() {
	fmt.Println(toLowerCase("Hello") == "hello")
	fmt.Println(toLowerCase("here") == "here")
	fmt.Println(toLowerCase("LOVELY") == "lovely")
}
func toLowerCase(str string) string {
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] <= 'Z' && bytes[i] >= 'A' {
			bytes[i] += 32
		}
	}
	return string(bytes)
}
