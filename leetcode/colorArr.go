package main

//颜色分球
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}

func sortArea(nums []int, target int) (counTaget int) {
	for i, num := range nums {
		if num == target {
			nums[i], nums[counTaget] = nums[counTaget], nums[i]
			counTaget++
		}
	}
	return
}

func sortColors(nums []int) []int {
	sortArea(nums[:sortArea(nums, 0)], 1)
	return nums
}
