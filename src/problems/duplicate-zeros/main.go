package main

import "fmt"

func main() {
	b := []int{1, 2, 0, 4, 5}
	duplicateZeros(b)
	fmt.Print(b)
}

func duplicateZeros(arr []int) {
	lenOfArr := len(arr)
	for i := 0; i < lenOfArr-1; i++ {
		if arr[i] == 0 {

			// Shift the array by onces to the right
			for j := lenOfArr - 1; j >= (i + 1); j-- {
				arr[j] = arr[j-1]
			}

			arr[i+1] = 0
			i++
		}
	}
}
