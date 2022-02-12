package main

import (
	"fmt"
)

func main() {
	fmt.Println("创建struct 的一些方式")

	s1 := new(chip)
	s1.id = 1
	(*s1).id = 11 // 也是合法的
	s1.name = "s1"
	s1.addId()

	s2 := &chip{id: 2, name: "s2"}
	s2.addId()

	s3 := chip{}
	s3.id = 3
	s3.name = "s3"
	s3.addId()

	s4 := newChip(4, "s4")
	s4.addId()
	fmt.Println(s1, s2, s3, s4)
}

type chip struct {
	id   int
	name string
}

func (c *chip) addId() {
	c.id += 100
}

func newChip(id int, name string) *chip {
	return &chip{id, name}
}