package main

import "fmt"

/*
* 功能：该函数用来实现插值查找算法
* 输入：查找数组values,数组长度n,查找元素element
* 输出：返回元素的位置
*/
func insertSearch(nums []int, element int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + (end - start)) * (element - nums[start]) / (nums[end] - nums[start])
		fmt.Printf("mid is :  %d \n", mid)
		if nums[mid] < element {
			start = mid + 1
		} else if nums[mid] > element {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	nums := []int{1, 7, 10, 13, 14, 15, 17, 20}
	fmt.Println(insertSearch(nums, 10))
	//fmt.Println(insertSearch(nums, 9))
	//fmt.Println(insertSearch(nums, 0))
}
