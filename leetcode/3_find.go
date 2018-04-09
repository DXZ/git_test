package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 0
	max_tmp := 0
	index_tmp := 0
	old_a := make(map[int]int, len(s))

	// aaa := [128]int{}
	for index, a := range s {
		// fmt.Println(aaa[a])
		// fmt.Println(int(a), max, max_tmp, index_tmp)
		if value, ok := old_a[int(a)]; ok && value >= index_tmp {
			if max_tmp > max {
				max = max_tmp
			}
			index_tmp = value + 1
			// fmt.Println("11111", index_tmp, value)
			max_tmp = index - index_tmp
			// fmt.Println("22222", max_tmp)
		}
		old_a[int(a)] = index
		max_tmp += 1
	}
	if max_tmp > max {
		max = max_tmp
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}
