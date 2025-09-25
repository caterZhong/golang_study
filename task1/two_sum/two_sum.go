package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var hashmap map[int]int = make(map[int]int, len(nums))
	for i, v := range nums {
		value, exists := hashmap[target-v]
		if exists {
			return []int{value, i}
		}
		hashmap[v] = i

	}

	return []int{}

}

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}
