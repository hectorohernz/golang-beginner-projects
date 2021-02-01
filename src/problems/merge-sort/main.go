package main

import "fmt"

func main() {
	b := []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	a := []int{6, 7, 8, 9, 10}
	fmt.Print(b, a)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	for n > 0 {
		nums1[m+n-1] = nums2[n-1]
		n--
	}
}
