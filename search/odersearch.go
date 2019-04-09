package main

import "fmt"

func orderSearch(nums []int, num int) int {
	for i, v := range nums {
		if num == v {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 5, 3, 4, 0}
	fmt.Println(orderSearch(nums, 10))
	fmt.Println(orderSearch(nums, 9))
	fmt.Println(orderSearch(nums, 0))
}
