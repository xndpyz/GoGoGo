package main

import "fmt"

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(arr []int, target int) []int {
	hash := make(map[int]int)
	for idx, val := range arr {
		if v, ok := hash[target-val]; ok {
			return []int{v, idx}
		}
		hash[val] = idx
	}
	return nil
}
