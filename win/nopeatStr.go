package main

import "fmt"

func main() {
	//最长不重复字符串
	noPeatStr("abcdabh")

}

func noPeatStr(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	left, right := 0, 1
	max := 0
	hash := map[byte]bool{}
	hash[s[0]] = true
	for right < len(s) {
		next := s[right]
		for hash[next] {
			delete(hash, s[left]) //保持最左边没有重复的字符串
			left++
		}
		if right-left+1 > max {
			max = right - left + 1
		}
		hash[next] = true
		right++
	}
	fmt.Println(max)
	return max
}
