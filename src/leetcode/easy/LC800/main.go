package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(similarRGB("#09f166") == "#11ee66")
}
func similarRGB(color string) string {
	var sb strings.Builder
	sb.Grow(7)
	sb.WriteByte('#')
	for i := 1; i < len(color); i += 2 {
		parseInt, _ := strconv.ParseInt(color[i:i+2], 16, 16)
		index := int(parseInt) / 17
		if parseInt%17 > 8 {
			index++
		}
		if index > 9 {
			v := byte('a' + index - 10)
			sb.WriteByte(v)
			sb.WriteByte(v)
		} else {
			v := strconv.Itoa(index)
			sb.WriteString(v)
			sb.WriteString(v)
		}
	}
	fmt.Println(sb.String())
	return sb.String()
}
