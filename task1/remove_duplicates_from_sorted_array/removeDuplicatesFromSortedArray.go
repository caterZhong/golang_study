package main

import "fmt"

// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
// 返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// leetcode: https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	var hashmap map[int]int = map[int]int{}
	originNums := make([]int, len(nums))
	copy(originNums, nums)

	index := 0
	for _, value := range originNums {
		_, exist := hashmap[value]
		length := len(nums)
		// 值存在， 说明重复， 则删除
		if exist {
			if index == length - -1 {
				nums = nums[0:len(hashmap)]
				return len(hashmap)
			} else {
				nums = append(nums[0:index], nums[index+1:]...)
			}
		} else {
			index++
			hashmap[value] = value
		}
	}
	return len(hashmap)

}

// 双指针解法
func removeDuplicatesWith2Pointer(nums []int) int {
	var hashmap map[int]int = map[int]int{}
	length := len(nums)
	desP := 0
	for searchP := 0; searchP < length; searchP++ {
		_, exist := hashmap[nums[searchP]]
		// 存在， 搜索指针继续往前， 目标指针不变
		if !exist {
			hashmap[nums[searchP]] = 1
			nums[desP] = nums[searchP]
			desP++
		}
	}
	return desP

}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicatesWith2Pointer(nums))
}
