package main

import (
	"math"
)

// 买卖股票的最佳时机
//输入：[7,1,5,3,6,4]
//输出：5
//
//输入：prices = [7,6,4,3,1]
//输出：0

//输入: prices = [7,1,5,3,6,4]
//输出: 7

func main() {
	twoProfit([]int{7, 1, 5, 3, 6, 4})
}

func twoProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			dp[i] = dp[i-1] + prices[i] - prices[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(prices)-1]
}

func maxProfit(prices []int) int {
	minValue := math.MinInt64
	maxValue := 0
	for i := 0; i < len(prices); i++ {
		if minValue > prices[i] {
			minValue = prices[i]
		} else if prices[i]-minValue > maxValue {
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}
