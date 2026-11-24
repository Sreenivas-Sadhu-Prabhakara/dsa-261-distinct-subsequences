package main

import "testing"

func TestNumDistinct(t *testing.T) {
	cases := []struct {
		s, t string
		want int
	}{
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
		{"abc", "abc", 1},
		{"abc", "abd", 0},
		{"aaa", "a", 3},
		{"anything", "", 1}, // the empty subsequence
		{"", "a", 0},
		{"", "", 1},
	}
	for _, c := range cases {
		if got := NumDistinct(c.s, c.t); got != c.want {
			t.Errorf("NumDistinct(%q,%q) = %d, want %d", c.s, c.t, got, c.want)
		}
	}
}
