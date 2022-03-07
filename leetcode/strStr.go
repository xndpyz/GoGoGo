package main

import "fmt"

//输入：haystack = "hello", needle = "ll"
//输出：2
func main() {
	//暴力破解
	//strStrs()

	// 滑动窗口法
	strStr("", "")
}

func strStr(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)
	if h < n {
		return -1
	}
	for i := n; i < h+1; i++ {
		if haystack[i-n:i] == needle {
			fmt.Println(i - n)
			return i - n
		}
	}
	return -1
}

func strStrs(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
outer:
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}
