package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}) == 1)
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}) == 2)
}
func numUniqueEmails(emails []string) int {
	m := map[string]bool{}
	for _, em := range emails {
		var sb strings.Builder
		for _, e := range em {
			if e == '.' {

			} else if e == '+' || e == '@' {
				sb.WriteString(em[strings.IndexByte(em, '@'):])
				break
			} else {
				sb.WriteByte(byte(e))
			}
		}
		m[sb.String()] = true
	}
	return len(m)
}
