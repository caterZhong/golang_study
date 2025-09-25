package main

import "fmt"

func get_only_one_num(nums []int) int {
	hashmap := map[int]int{}
	for _, value := range nums {
		count := hashmap[value]
		count++
		hashmap[value] = count
	}

	for key, value := range hashmap {
		if value == 1 {
			return key
		}
	}

	return 0
}

func main() {
	nums := []int{3, 2, 4, 2, 4, 3, 5}
	fmt.Println("只出现一次的数字", get_only_one_num(nums))
}
