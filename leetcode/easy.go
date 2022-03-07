package main

import (
	"fmt"
	"sort"
)

func main() {
	res := search([]int{1, 2, 3, 4, 6, 8}, 5)
	fmt.Println(res)

	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 6)

	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})

	twoSumNumber([]int{1, 2, 3, 4}, 5)
}

func twoSumNumber(nums []int, target int) []int {
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if value, ok := hash[target-nums[i]]; ok {
			return []int{value, i}
		}
		hash[nums[i]] = i
	}
	return nil
}

//[6,2,6,5,1,2]  1,2,2,5,6,6  [1,2,3,4] 1,2, 3,4
func arrayPairSum(nums []int) (ans int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]

		right--
		left++
	}
}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		fmt.Println((i + k) % len(nums))
		newNums[(i+k)%len(nums)] = v

	}
	copy(nums, newNums)
	fmt.Println(nums)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	var insert int
	for left < right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
			insert = left
		} else {
			right = mid
			insert = right
		}
	}
	fmt.Println(insert)
	return -1
}

func searchInsert(nums []int, target int) int {

	return 1
}
