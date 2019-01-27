package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) < 1 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else {
		tag := -1
		flag := false
		for index := range strs[0] {
			if flag {
				break
			}
			for j := range strs {
				if len(strs[j])-1 < index || strs[0][index] != strs[j][index] {
					tag = index
					flag = true
					break
				}
			}
		}
		if tag == -1 {
			return strs[0]
		}
		return string([]rune(strs[0])[:tag])
	}
}

func main() {
	// fmt.Println(longestCommonPrefix([]string{"cc", "c"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
