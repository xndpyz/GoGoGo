package main

import (
	"fmt"
)

func main() {
	//斐波那契额数列
	//feibonaacci(4)

	//爬楼梯
	climbStairs(5)

	//**爬楼梯的最小费用
	minCostClimbingStairs([]int{10, 15, 10})

	//打家劫舍 状态转移
	ponch([]int{3, 2, 3, 48, 7})

	//打家劫舍升级版 房屋围成了一圈， 第一个和最后一个是相邻的
	//切片分成两部分来操作即可 [:ln-1], [1:]

	//获取最大点数 转化成打劫家舍来做的
	deleteAndEarn([]int{2, 2, 3, 3, 3, 4})

	//跳跃游戏
	jumpGame([]int{2, 3, 1, 1, 4})
	jumpGame([]int{3, 2, 1, 0, 4})

	//最大子数组和
	maxChirldSum([]int{-2})

	//环形数组最大和
	maxRangeSum([]int{5, -1, 5})

	//杨辉三角
	yangHuiSanjiao(3)

	//最大子数组乘积和, 并不能用动态规划
	maxXSum([]int{2, 3, 3, 4, -2, 4})
}

func maxXSum(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	sign := true
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 && sign == true {
			dp[i] = nums[i] * dp[i-1]
		} else {
			dp[i] = max(nums[i], dp[i-1])
			sign = false
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

func yangHuiSanjiao(rowIndex int) []int {
	dp := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp[rowIndex-1]
}

func maxRangeSum(nums []int) int {

	return 0
}

func maxChirldSum(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		res = max(res, dp[i])
	}
	return res
}

func jumpGame(nums []int) bool {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < i {
			return false
		}
		dp[i] = max(dp[i-1], i+nums[i])
	}
	return true
}

func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, v := range nums {
		maxVal = max(v, maxVal)
	}
	arr := make([]int, maxVal+1)
	for _, v := range nums {
		arr[v] += v //这里等累加，计算有重复的数字
	}
	//sort.Ints(arr)
	//fmt.Println(arr)
	dp := make([]int, len(arr))
	dp[0], dp[1] = arr[0], max(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i])
	}
	//fmt.Println(dp)
	return dp[len(dp)-1]
}

func ponch(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}
	if ln == 1 {
		return nums[0]
	}
	dp := make([]int, ln)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < ln; i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	fmt.Println(dp)
	return dp[ln-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minCostClimbingStairs(cost []int) int {
	ln := len(cost)
	dp := make([]int, ln+1)
	for i := 2; i < ln+1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	fmt.Println(dp[ln])
	return dp[ln]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func climbStairs(n int) int {
	hash := map[int]int{
		1: 1, 2: 2,
	}
	for i := 3; i < n+1; i++ {
		hash[i] = hash[i-1] + hash[i-2]
	}
	fmt.Println(hash)
	return hash[n]
}

func feibonaacci(num int) int {
	hash := map[int]int{
		1: 1, 2: 1,
	}
	for i := 3; i < num+1; i++ {
		hash[i] = hash[i-1] + hash[i-2]
	}
	fmt.Println(hash)
	return hash[num]
}
