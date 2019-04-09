package main

import "fmt"

/*
704. Binary Search

Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.


Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
 */

func search(nums []int, target int) int {
	low,high:=0,len(nums)-1
	for low <=high{
		mid := (low+high)/2
		if nums[mid] <target {
			low =mid+1
		}else if nums[mid] > target {
			high=mid -1
		}else {
			return mid
		}
	}
	return -1
}


func main() {
	nums := []int {-1,0,3,5,9,12}
	fmt.Println(search(nums,-1))
	nums =[]int {-1,0,3,5,9,12}
	fmt.Println(search(nums,2))
}