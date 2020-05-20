package 字符串匹配

import "testing"

func Test_RabinKarp(t *testing.T) {
	patString := "abcd"
	rk := NewInstanceRabinKarp(patString)
	tests := []struct {
		param string
		index int
	}{
		{"aabcde", 1},
		{"112345", -1},
		{"asfd oaisdufkhasdgfisdfghaskfyaeiurnfkfabcdefaksjfhksjdfhk", 39},
		{"abcde", 0},
		{"abcdee", 0},
		{"abcdevfddfgdsgdsfge", 0},
		{"bcdee", -1},
		{"asfd oaisdufkhasdgfisdfghaskfyaeiurnfkfabcde", 39},
	}
	for _, test := range tests {
		if got := rk.Search(test.param); got != test.index {
			t.Errorf("rk.Search(%q,%q) = %v", test.param, patString, got)
		}
		if got := Search(test.param, patString); got != test.index {
			t.Errorf("Search(%q,%q) = %v", test.param, patString, got)
		}
	}
}
