package main

import "fmt"

func main() {
	fmt.Println(isValid("(){}"))
	fmt.Println(isValid("[({})]"))
	fmt.Println(isValid("]{"))

}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	hash := map[byte]byte{
		'{': '}', '(': ')', '[': ']',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || s[i] != hash[stack[len(stack)-1]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
