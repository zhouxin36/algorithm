package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//["9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"]
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	//["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}
	for _, cp := range cpdomains {
		split := strings.Split(cp, " ")
		num, _ := strconv.Atoi(split[0])
		dom := strings.Split(split[1], ".")
		for i := 0; i < len(dom); i++ {
			m[strings.Join(dom[i:], ".")] += num
		}
	}
	var result []string
	for k, v := range m {
		result = append(result, strconv.Itoa(v)+" "+k)
	}
	return result
}
