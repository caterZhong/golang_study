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

	// for _, v := range digits {
	// 	num = num*10 + v
	// }

	// num++
	// var result []int = []int{}
	// var i int = 0
	// for {

	// 	digit := num % 10
	// 	num = num / 10
	// 	if len(result) == 0 {
	// 		result = append(result, digit)
	// 	} else {
	// 		result = append([]int{digit}, result[0:]...)
	// 	}

	// 	i++||~|
	// 	if num == 0 {
	// 		break
	// 	}

	// }
	// return result
}

func main() {
	nums := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(nums))
}
