package main

import (
	"fmt"
)

func main() {
	fmt.Println("sss:", longestPalindrome("caac"))
}

func longestPalindrome(s string) string {
	str := []byte(s)
	if len(str) <= 1 {
		return s
	}

	start, end := 1, len(str)
	maxlegth := 0
	var max string
	for start < end {
		begin1 := start
		begin0 := start
		if 2*(end-begin1) < maxlegth {
			break
		}
		for begin1 < end-1 && str[begin1] == str[begin1+1] {
			begin1++
		}
		fmt.Println("111", begin1, begin0)
		for begin0-1 > 0 && begin1+1 < end && str[begin0-1] == str[begin1+1] {
			fmt.Println("3333", begin0, begin1, string(str[begin0]), string(str[begin1]))
			begin0--
			begin1++
		}

		fmt.Println("222", begin1, begin0)
		if maxlegth < begin1-begin0 {
			maxlegth = begin1 - begin0

			max = string(str[begin0 : begin1+1])
			fmt.Println(begin1, begin0, max)
		}
		start++
	}

	return string(max)
}

// func longestPalindrome3(s string) (result string) {
// 	s1 := reverseString(s)
// 	runes1 := []rune(s1)
// 	runes := []rune(s + s)
// 	var result_tmp []rune
// 	var result_runes []rune
// 	for i := 0; i < len(runes); i++ {
// 		k := i

// 		for j := 0; j < len(runes1) && k < len(runes); j++ {
// 			if runes[k] == runes1[j] {
// 				fmt.Println("111", string(runes[k]), k, j)
// 				result_tmp = append(result_tmp, runes[k])
// 			}
// 			if runes[k] != runes1[j] || j == len(runes1)-1 || k+1 == len(runes1) {
// 				fmt.Println("222", string(result_tmp), k, j, len(result_tmp), len(runes1))
// 				if len(result_tmp) > len(result_runes) && (k+j-len(result_tmp))%len(runes1) == 0 {
// 					result_runes = result_tmp
// 				}
// 				result_tmp = []rune{}
// 			}
// 			k++
// 		}
// 		fmt.Println("111212312312", string(result_tmp))
// 		result_tmp = []rune{}
// 	}

// 	result = string(result_runes)
// 	return
// }

func longestPalindromeold(s string) (result string) {
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
				// fmt.Println("aaaa", k, j, string(result_tmp))
				if len(result_tmp) > len(result_runes) && (k+j-len(result_tmp))%len(runes1) == 0 {
					result_runes = result_tmp
				}
				result_tmp = []rune{}
			}
			if runes[k] == runes1[j] {
				// fmt.Println("111", string(runes[k]), k, j)
				result_tmp = append(result_tmp, runes[k])
				k++
				continue
			}
			// fmt.Println("bbbb", k, j, string(result_tmp), (k+j)%len(runes1), len(runes1))
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
