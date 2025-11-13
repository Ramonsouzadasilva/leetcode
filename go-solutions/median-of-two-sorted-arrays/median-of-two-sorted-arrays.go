package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		px := (low + high) / 2
		py := (x + y + 1) / 2 - px

		var maxX, minX, maxY, minY int

		if px == 0 {
			maxX = math.MinInt
		} else {
			maxX = nums1[px-1]
		}

		if px == x {
			minX = math.MaxInt
		} else {
			minX = nums1[px]
		}

		if py == 0 {
			maxY = math.MinInt
		} else {
			maxY = nums2[py-1]
		}

		if py == y {
			minY = math.MaxInt
		} else {
			minY = nums2[py]
		}

		if maxX <= minY && maxY <= minX {
			if (x+y)%2 == 0 {
				return float64(max(maxX, maxY)+min(minX, minY)) / 2.0
			}
			return float64(max(maxX, maxY))
		}

		if maxX > minY {
			high = px - 1
		} else {
			low = px + 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))        
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))    
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))    
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))           
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))           
}