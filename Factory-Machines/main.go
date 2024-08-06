package main

import "fmt"

func main() {

	// n -> machines , make t products
	// machines can work simultaneously -> shortest time needed to make t products
	var t int = 7

	machines := []int{3, 2, 5}

	low := 1
	high := int(1e18)

	var minTime int = high

	for low <= high {
		mid := low + (high - low) / 2
		if predicate(machines, t, mid) {
			minTime = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println("Min. time req. to make t products = ", minTime)

}

func predicate(machines []int, t int, timeTaken int) bool {
	products := 0
	for _, machineTime := range(machines) {
		products += timeTaken / machineTime
		if products >= t {
			return true
		}
	} 
	return products >= t
}