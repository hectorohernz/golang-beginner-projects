package main

import "fmt"

func main() {
	a := []int{7, 3, 43, 7, 2, 5, 1, 45, 6, 3}
	num := selectionSort(a)
	fmt.Println(num)
}

func selectionSort(arr []int) int {
	smallestArr := arr[0]
	smallestIdx := 0
	lenOfA := len(arr)
	for i := 1; i < lenOfA; i++ {
		if arr[i] < smallestArr {
			smallestArr = arr[i]
			smallestIdx = i
		}
	}
	return smallestIdx
}
