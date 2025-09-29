package main

import "fmt"

func addTen(num *int) {
	*num += 10
}

func mutiplyTwo(numsP *[]int) {
	for i, value := range *numsP {
		(*numsP)[i] = value * 2
	}
}

func main() {
	var num int = 5
	addTen(&num)
	fmt.Println("修改后的值：", num)

	var nums []int = []int{1, 2, 3}
	mutiplyTwo(&nums)
	fmt.Println("修改后slice的值：", nums)
}
