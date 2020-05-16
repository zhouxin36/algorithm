package 字符串匹配

import "testing"

func Test_RabinKarp(t *testing.T) {
	patString := "abcd"
	rk := NewInstanceRabinKarp(patString)
	tests := []struct {
		param string
		index int
	}{
		{"112345", -1},
		{"aabcde", 1},
		{"asfd oaisdufkhasdgfisdfghaskfyaeiurnfkfabcdefaksjfhksjdfhk", 39},
		{"abcde", 0},
		{"abcdee", 0},
		{"abcdevfddfgdsgdsfge", 0},
		{"bcdee", -1},
		{"asfd oaisdufkhasdgfisdfghaskfyaeiurnfkfabcde", 39},
	}
	for _, test := range tests {
		if got := rk.Search(test.param); got != test.index {
			t.Errorf("Search(%q) = %v", test.param, got)
		}
	}
}
