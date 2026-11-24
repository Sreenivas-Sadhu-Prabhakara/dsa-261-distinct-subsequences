package main

import "fmt"

// NumDistinct counts how many distinct subsequences of s spell exactly t.
// One row over t, swept RIGHT-TO-LEFT per character of s so each s-character
// is used at most once per subsequence.
func NumDistinct(s, t string) int {
	m := len(t)
	dp := make([]int, m+1) // dp[j]: ways to spell t[:j] from the s seen so far
	dp[0] = 1              // the empty target: one way, take nothing
	for i := 0; i < len(s); i++ {
		for j := m; j >= 1; j-- {
			if t[j-1] == s[i] {
				dp[j] += dp[j-1] // use s[i] as t's j-th character, or don't (dp[j] keeps that)
			}
		}
	}
	return dp[m]
}

func main() {
	fmt.Println(NumDistinct("rabbbit", "rabbit"))
	fmt.Println(NumDistinct("babgbag", "bag"))
}
