package main

import (
		"fmt"
		)

/**
求俩个数组的交集  349
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Each element in the result must be unique.
The result can be in any order.

 */

/*
先求出他们的并集，
然后对并集进行排序
 */
func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)

	for i:=0;i<len(nums1);i++ {
		m1[nums1[i]]=true
	}
	reset := make(map[int]bool)
	internums := []int{}
	for _,k :=range nums2 {
		if m1[k] {
			reset[k]=true
		}
	}

	for k,v :=range reset {
		if v {
			internums =append(internums,k)
		}
	}

	for i:=0;i< len(internums);i++ {
		for j:=0;j<len(internums)-1 -i;j++ {
			if internums[j] > internums[j+1] {
				internums[j],internums[j+1] = internums[j+1],internums[j]
			}
		}
	}

	return internums

}

func main() {
	nums1 := []int {1,2,2,1}
	nums2 := []int {2,2}
	res :=intersection(nums1,nums2)

	for i:=0;i<len(res);i++{
		fmt.Println(res[i])
	}
	fmt.Println("----------")

	nums1 = []int {4,9,5}
	nums2 = []int {4,9,9,8,4,5}
	res = intersection(nums1,nums2)

	for i:=0;i<len(res);i++{
		fmt.Println(res[i])
	}
}
