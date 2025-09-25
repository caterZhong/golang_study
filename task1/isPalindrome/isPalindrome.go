package main

import (
	"fmt"
	"strconv"
)

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// leetcode ： https://leetcode.cn/problems/palindrome-number/description/

func isPalindrome(x int) bool {
	strx := strconv.Itoa(x)
	length := len(strx)
	for i, j := 0, length-1; i <= j; i, j = i+1, j-1 {
		if strx[i] != strx[j] {
			return false
		}
	}
	return true

}

// func isPalindrome(x int) bool {
// 	strx := strconv.Itoa(x)
// 	if

// }

func main() {
	x := 1211211
	fmt.Println("x 回文数判定结果：", isPalindrome(x))
}
