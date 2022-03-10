package main

import "fmt"

//容量= len(原容量-左值）长度=len(右边-左边）

func main() {
	//fmt.Println("切片相关")
	//var a1 = []int{0, 0, 0}
	//fmt.Println(a1)

	//a := make([]int, 0, 4) //长度为0， 容量为4
	//a := []int{1, 2, 3}
	//b := append(a, 10)
	//c := append(a, 20, 20, 3, 4, 40)
	//d := append(a, 12, 2, 3, 4, 5)
	//fmt.Println(a, b, c, d)
	//fmt.Printf("%p, %p, %p\n", b, c, d)

	//q1()
	//q2

	q3()
}

func q3() {
	a := []int{1, 2}
	b := append(a, 3)
	c := append(b, 4)
	d := append(b, 5)
	e := append(b, 6, 7)
	g := append(b, 8, 9, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println(a, b, c, d, e, g)
	fmt.Println(e, len(e), cap(e))
	fmt.Println(g, len(g), cap(g))
	fmt.Printf("%p %p %p %p", a, b, c, d)
}

func q1() {
	a := []int{0, 1, 2, 3}
	x := a[:1]          //[0], len=1, cap=4
	y := a[2:]          //[2, 3], len=2, cap=2
	x = append(x, y...) // x [0, 2, 3] a [0, 2, 3, 3] [3,3]
	fmt.Printf("%p, %p, %p\n", a, x, y)
	x = append(x, y...) //
	fmt.Println(a, x)
	//fmt.Printf("%p, %p, %p\n", a, x, y)

	s := []int{1, 2}
	s = append(s, 4, 5)
	fmt.Println(len(s), cap(s)) //5, 6
}

func q2() {
	a := make([]int, 0, 5)
	b := a[:]
	//c := a[1:]  //编译报错
	d := a[1:4]
	e := a[:4]
	g := a[2:4]
	fmt.Println("b:", len(b), cap(b))
	//fmt.Println("c:", len(c), cap(c))
	fmt.Println("d:", len(d), cap(d))
	fmt.Println("e:", len(e), cap(e))
	fmt.Println("g:", len(g), cap(g))

	s := make([]int, 5, 5)
	s2 := s[2:4]
	fmt.Println(s2, len(s2), cap(s2))
}

func changeSlice(arr []int) {
	arr[1] = 1
}
