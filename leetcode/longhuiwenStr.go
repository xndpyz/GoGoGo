package main

//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//示例 2：
//
//输入：s = "cbbd"
//输出："bb"

func main() {
	longestPalindrome("aabaa")
}

func longestPalindrome(s string) string {
	// 双指针
	start, end := 0, 0
	ln := len(s)
	for i := 0; i < ln; {
		left, right := i, i
		for right < ln-1 && s[right] == s[right+1] {
			right++
		}
		i = right + 1
		for left > 0 && right < ln-1 && s[left-1] == s[right+1] {
			left--
			right++
		}
		if end-start < right-left {
			start = left
			end = right
		}
	}
	return s[start : end+1]
}