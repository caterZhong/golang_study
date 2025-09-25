package main

import "fmt"

func plusOne(digits []int) []int {
	var plusone bool = true
	var length int = len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] == 9 && plusone {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			return digits
		}
	}
	return append([]int{1}, digits[0:]...)

}

func main() {
	nums := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(nums))
}
