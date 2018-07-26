package main

import (
	"fmt"
	"reflect"
)

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	if len(s) == 0 && len(p) == 1 {
		return false
	}

	if len(s) == 0 && len(p) >= 2 && string(p[1]) != "*" {
		return false
	}
	if len(s) == 0 && len(p) >= 2 && string(p[1]) == "*" {
		return isMatch(s, p[2:])
	}

	if len(p) == 0 && len(s) != 0 {
		return false
	}

	fmt.Println(p, "----", s)

	if p[0:1] != "." && p[0:1] != s[0:1] && (len(p) < 2 || (len(p) >= 2 && p[1:2] != "*")) || (p[len(p)-1:] != s[len(s)-1:] && p[len(p)-1:] != "*" && p[len(p)-1:] != ".") {
		return false
	}

	if len(p) >= 2 && p[1:2] == "*" {
		if p[0:1] != s[0:1] && p[0:1] != "." {
			return isMatch(s[:], p[2:])
		} else {
			if isMatch(s[1:], p[2:]) {
				return true
			} else if isMatch(s[:], p[2:]) {
				return true
			} else {
				return isMatch(s[1:], p[:])
			}
		}

	} else {
		return isMatch(s[1:], p[1:])
	}

}

func isMatch2(s string, p string) bool {
	slen, plen := len(s), len(p)
	var dp [][]bool
	var t []bool
	for i := 0; i <= slen; i++ {
		t = make([]bool, plen+1)
		dp = append(dp, t)
	}
	for i := 0; i <= slen; i++ {
		for j := 0; j <= plen; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			} else if i == 0 {
				dp[i][j] = ((j-1)%2 == 1 && p[j-1] == '*' && dp[i][j-2])
				continue
			} else if j == 0 {
				dp[i][j] = false
				continue
			}
			if p[j-1] != '*' {
				dp[i][j] = (p[j-1] == s[i-1] || p[j-1] == '.') && dp[i-1][j-1]
			} else {
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i-1][j]
				}
				if dp[i][j-2] == true {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[slen][plen]
}

func main() {
	s := "asd"
	fmt.Println(reflect.TypeOf(string(s[1])), reflect.TypeOf(s[2:3]), s[2:3], s[len(s)-1:])
	fmt.Println(isMatch("mississippi", "a*a*a*a*a*a*a*a*a*a*c"))
}

// "aaaaaaaaaaaaab"
// "a*a*a*a*a*a*a*a*a*a*c"
// "mississippi"
// "mis*is*p*."
