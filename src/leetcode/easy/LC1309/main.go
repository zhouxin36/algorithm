package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(freqAlphabets("1326#") == "acz")
	fmt.Println(freqAlphabets("10#11#12") == "jkab")
	fmt.Println(freqAlphabets("25#") == "y")
	fmt.Println(freqAlphabets("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#") == "abcdefghijklmnopqrstuvwxyz")
}
func freqAlphabets(s string) string {
	split := strings.Split(s, "#")
	var res strings.Builder
	for i := 0; i < len(split)-1; i++ {
		if len(split[i]) > 2 {
			for _, v := range split[i][:len(split[i])-2] {
				res.WriteByte(byte(v - '0' + 'a' - 1))
			}
			v, _ := strconv.Atoi(split[i][len(split[i])-2:])
			res.WriteByte(byte(v + 'a' - 1))
		} else {
			v, _ := strconv.Atoi(split[i])
			res.WriteByte(byte(v + 'a' - 1))
		}
	}
	if split[len(split)-1] != "" {
		for _, v := range split[len(split)-1] {
			res.WriteByte(byte(v - '0' + 'a' - 1))
		}
	}
	return res.String()
}
