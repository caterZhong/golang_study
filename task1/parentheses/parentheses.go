package main

import "fmt"

func isValid(s string) bool {

	var stack []byte = []byte{}

	var length int = len(s)

	if length == 0 {
		return true
	}

	if length%2 != 0 {
		return false
	}

	for i := 0; i < length; i = i + 1 {
		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
			stack = append(stack, s[i])
		}

		if s[i] == ']' || s[i] == '}' || s[i] == ')' {
			topIndex := len(stack) - 1
			if topIndex < 0 {
				return false
			}
			if !isPairs(stack[topIndex], s[i]) {
				return false
			}
			stack = stack[:topIndex]
		}
	}
	return len(stack) == 0

}

func isPairs(a byte, b byte) bool {
	if a == '{' && b == '}' {
		return true
	}

	if a == '(' && b == ')' {
		return true
	}

	if a == '[' && b == ']' {
		return true
	}
	return false
}

func main() {
	s := "))"
	fmt.Println("括号是否匹配", isValid(s))
}
