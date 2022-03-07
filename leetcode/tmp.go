package main

import "fmt"

func main() {
	// 求最长不重复字符串
	f1("abcabd")

	// 求最长回文字符串
	f2("aabbaa")

	f3([]int{1, 2, 3, 4, 5}, 6)

	//判断回文数51115==> 511, 51
	f4(51115)
}

func f4(num int) bool {
	if num < 0 || num%10 == 0 {
		return false
	}
	n := 0
	for num > n {
		n = n*10 + num%10
		num /= 10
		fmt.Println(n, num)
	}
	fmt.Println(n, num, n/10)
	return num == n || n/10 == num
}

func f3(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = nums[mid]
		} else {
			left = nums[mid] + 1
		}
	}
	return -1
}

func f2(s string) string {
	ln := len(s)
	start, end := 0, 0
	for i := 0; i < ln; i++ {
		right, left := i, i
		for right < ln-1 && s[right] == s[right+1] {
			right++
		}
		//right++
		for left > 0 && right < ln-1 && s[left-1] == s[right+1] {
			right++
			left--
		}
		if end-start < right-left {
			end = right
			start = left
		}
	}
	fmt.Println(end, start)
	return ""
}

func f1(s string) int {
	hash := map[byte]int{}
	right, res := 0, 0
	// 滑动窗口
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(hash, s[i-1])
		}
		for right < len(s) && hash[s[right]] == 0 {
			hash[s[right]]++
			right++
		}
		res = Max(res, right-i)
	}
	fmt.Println("最小长度为：", res)
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
