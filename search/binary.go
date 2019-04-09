package main

import (
	"fmt"
	"math"
)

/**
二分查找
 */

func binarySearch(nums []int, num int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums) -1
	mid :=int(math.Floor(float64((low + high) / 2)))
	for low <= high {
		if num < nums[mid] {
			high = mid -1
		} else if num > nums[mid] {
			low = mid +1
		} else {
			return mid

		}
		mid = int(math.Floor(float64((low + high) / 2)))
	}
	return -1

}



func main() {
	nums := []int{0,1,2,3,4,5,7,10}
	fmt.Println(binarySearch(nums, 10))
	fmt.Println(binarySearch(nums, 9))
	fmt.Println(binarySearch(nums, 0))
}
