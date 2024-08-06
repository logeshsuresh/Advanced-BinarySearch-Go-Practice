package main 

import (
	"fmt"
)

func main() {
	array := []int32{2, 4, 7, 3, 5}
	var k int32 = 3

	low, high := findMax(array), findSum(array)	

	var ans int32 = high

	for low <= high {
		mid := low + (high - low) / 2
		if predicate(array, k, mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println("Maximum sum in a sub-array = ", ans)
}

func findMax(array []int32) int32 {
	max := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func findSum(array[] int32) int32 {
	var sum int32
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func predicate(array []int32, k int32, maxSum int32) bool {
	var curSum int32 = 0
	var curArrays int32 = 1
	for i := 0; i < len(array); i++ {
		if array[i] > maxSum {
			return false
		} else if curSum + array[i] <= maxSum {
			curSum += array[i]
		} else {
			curSum = array[i]
			curArrays++
		}
	}
	return curArrays <= k
}