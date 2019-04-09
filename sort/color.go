package main

import "fmt"

/*
sort colors

Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

采用技术排序来做
 */

func sortColors(nums []int) {
	temp := make([]int, 3)
	for _, v := range nums {
		temp[v] ++
	}

	idx := 0
	for i, v := range temp {
		for j := 0; j < v; j++ {
			nums[idx] = i
			idx++
		}
	}

}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	for _, v := range nums {
		fmt.Println(v)
	}
}
