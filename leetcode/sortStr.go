package main

import "fmt"

func main() {
	sortStr("ab", "eidbaooo")
}

//判断s1 是否为s2的字串, 双指针
//输入：s1 = "ab" s2 = "eidbaooo"
//输出：true
//解释：s2 包含 s1 的排列之一 ("ba").
func sortStr(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	fmt.Println(cnt)
	left := 0
	for right, ch := range s2 {
		x := ch - 'a' //字母索引
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		fmt.Println(right - left + 1)
		if right-left+1 == n {
			return true
		}
	}
	return false
}
