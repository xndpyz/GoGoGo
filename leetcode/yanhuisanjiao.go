package main

import "fmt"

//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]] [1, 5, 10, 10, 5, 1]
func main() {
	fmt.Println("杨辉三角")
	generateSanjiao(5)
}

func generateSanjiao(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	fmt.Println(res)
	return res
}
