package main

import (
	"fmt"
	"math"
)

/*
判断一个num是否满足正方形所需条件
 */
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	low, high := 1, num
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}


func main() {
	fmt.Println(math.Sqrt(104976))
	fmt.Println(isPerfectSquare(104976))
	fmt.Println(isPerfectSquare(1))
}