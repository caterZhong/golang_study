package main

import "fmt"

func merge(intervals [][]int) [][]int {
	selectedSort(intervals)
	for i, j := 0, 1; j < len(intervals); {
		intervalI := intervals[i]
		intervalJ := intervals[j]
		// 不重叠
		if intervalI[1] < intervalJ[0] {
			i++
			j++
		} else {
			maxR := intervalJ[1]
			if intervalI[1] > intervalJ[1] {
				maxR = intervalI[1]
			}
			newInterval := [][]int{{intervalI[0], maxR}}
			intervals = append(append(intervals[0:i], newInterval...), intervals[j+1:]...)
		}
	}
	return intervals
}

func selectedSort(intervals [][]int) {
	length := len(intervals)
	for i := 0; i < length; i++ {
		minL := intervals[i][0]
		minIndex := i
		for j := i + 1; j < length; j++ {
			if intervals[j][0] < minL {
				minL = intervals[j][0]
				minIndex = j
			}
		}
		// 找到本轮最小的，交换
		if i != minIndex {
			temp := intervals[i]
			intervals[i] = intervals[minIndex]
			intervals[minIndex] = temp
		}
	}
}

func main() {

	intervals := [][]int{{8, 10}, {1, 3}, {2, 6}, {15, 18}}
	fmt.Println(merge(intervals))
	// fmt.Print	ln(intervals[:])
	// fmt.Println(intervals[3:])

	// i2 < j1
}
