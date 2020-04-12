package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(int32(1e9 + 7))
	//& is an HTML entity but &ambassador; is not.
	fmt.Println(entityParser("&amp; is an HTML entity but &ambassador; is not."))
	//and I quote: \"...\"
	fmt.Println(entityParser("and I quote: &quot;...&quot;"))
}
func entityParser(text string) string {
	v := text
	v = strings.ReplaceAll(v, "&quot;", "\"")
	v = strings.ReplaceAll(v, "&apos;", "'")
	v = strings.ReplaceAll(v, "&amp;", "&")
	v = strings.ReplaceAll(v, "&gt;", ">")
	v = strings.ReplaceAll(v, "&lt;", "<")
	v = strings.ReplaceAll(v, "&frasl;", "/")
	return v
}
