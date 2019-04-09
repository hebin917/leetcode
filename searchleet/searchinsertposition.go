package main

import "fmt"

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0

 */
func searchInsert(nums []int, target int) int {
	//如果匹配到，则返回匹配的index
	//如果没有匹配到，则找到插入点：插入点需要满足，i < nums[mid] < i+1
	if len(nums) ==1 {
		if nums[0] ==target || nums[0] > target{
			return 0
		}else {
				return 1
		}
	}

	lo,hi :=0,len(nums)-1
	//find :=true
	for lo <= hi {
		mid := (lo + hi)/2
		if nums[mid] ==target {
			return mid
		}else if nums[mid] > target {
			hi = mid -1
		}else {
			lo = mid +1
		}
	}
	return lo
}


func main() {
	nums := []int {1,3,5,6}
	fmt.Println(searchInsert(nums,5))
	fmt.Println(searchInsert(nums,2))
	fmt.Println(searchInsert(nums,7))
	fmt.Println(searchInsert(nums,0))
}


