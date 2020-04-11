package main

import "strings"

func main() {

}
func toGoatLatin(S string) string {
	m := map[byte]struct{}{}
	for _, c := range []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
		m[c] = struct{}{}
	}
	var sb strings.Builder
	t := 1
	for _, w := range strings.Split(S, " ") {
		first := w[0]
		if _, ok := m[first]; ok {
			sb.WriteString(w)
		} else {
			sb.WriteString(w[1:])
			sb.WriteByte(w[0])
		}
		sb.WriteString("ma")
		for i := 0; i < t; i++ {
			sb.WriteByte('a')
		}
		t++
		sb.WriteByte(' ')
	}
	return sb.String()[:sb.Len()-1]
}
