package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 10, 11, 22, 42, 53}
	fmt.Println(binarySearch(a, 0))
}

func binarySearch(arr []int, target int) int {
	l := len(arr)
	// Edge Case
	if l == 1 && arr[0] == target {
		return 0
	}

	min := 0
	max := l - 1
	for max > 0 {
		average := (min + max) / 2
		if arr[average] == target {
			return average
		}
		if arr[average] < target {
			min = average + 1
		} else {
			max = average - 1
		}
	}
	return -1
}
