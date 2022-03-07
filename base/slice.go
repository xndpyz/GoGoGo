package main

import "fmt"

func main() {
	fmt.Println("切片相关")

	var a1 = []int{0, 0, 0}
	changeSlice(a1)
	fmt.Println(a1)
}

func changeSlice(arr []int) {
	arr[1] = 1
}
