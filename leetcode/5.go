package main

import (
	"fmt"
)

func longestPalindrome(s string) string {

	// for i in
	return ""
}

func longestPalindrome3(s string) (result string) {
	// l1 := len(s)
	s1 := reverseString(s)
	// fmt.Println(l1, s1)
	runes1 := []rune(s1)
	runes := []rune(s + s)
	// result = result
	fmt.Println(s, s1)
	var result_tmp []rune
	var result_runes []rune
	for i := 0; i < len(runes); i++ {
		k := i

		for j := 0; j < len(runes1) && k < len(runes); j++ {
			if len(runes1) == k {
				fmt.Println("aaaa", k, j, string(result_tmp))
				if len(result_tmp) > len(result_runes) && (k+j-len(result_tmp))%len(runes1) == 0 {
					result_runes = result_tmp
				}
				result_tmp = []rune{}
			}
			if runes[k] == runes1[j] {
				fmt.Println("111", string(runes[k]), k, j)
				result_tmp = append(result_tmp, runes[k])
				k++
				continue
			}
			fmt.Println("bbbb", k, j, string(result_tmp), (k+j)%len(runes1), len(runes1))
			if len(result_tmp) > len(result_runes) && (k+j-len(result_tmp))%len(runes1) == 0 {
				result_runes = result_tmp
			}
			result_tmp = []rune{}
			k++

		}

		if len(result_tmp) > len(result_runes) && (len(runes1)+k-len(result_tmp))%len(runes1) == 0 {
			result_runes = result_tmp
		}
		result_tmp = []rune{}
	}
	if len(result_tmp) > len(result_runes) {
		result_runes = result_tmp
		result_tmp = []rune{}
	}
	result = string(result_runes)
	return
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func main() {
	fmt.Println("sss:", longestPalindrome3("abcdasdfghjkldcba"))
}
