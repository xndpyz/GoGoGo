package main

import "fmt"

//例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
func main() {
	lengthOfLongestSubstring("abcabcbb")
}

func lengthOfLongestSubstring(s string) int {
	hash := map[byte]int{}
	right, res := -1, 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(hash, s[i-1])
		}
		for right+1 < len(s) && hash[s[right+1]] == 0 {
			hash[s[right+1]]--
			right++
		}
		res = maxNum(res, right+1-i)
	}
	fmt.Println(res)
	return res
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
