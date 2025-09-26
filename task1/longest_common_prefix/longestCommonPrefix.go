package main

import "fmt"

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	for _, v := range strs {
		if len(v) < minLen {
			minLen = len(v)
		}
	}

	var longestCommonPrefix string = ""
	for i := 0; i < minLen; i++ {
		var ch byte = strs[0][i]
		for _, v := range strs {
			if v[i] != ch {
				return longestCommonPrefix
			}
		}
		longestCommonPrefix += string(ch)
	}

	return longestCommonPrefix

}

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
