package main

import (
	"strconv"
	"sort"
	"strings"
	"fmt"
)

/**
179. Largest Number

Given a list of non negative integers, arrange them such that they form the largest number.

Example 1:

Input: [10,2]
Output: "210"
Example 2:

Input: [3,30,34,5,9]
Output: "9534330"



 */

type ByLength []string

func (a ByLength) Len() int {
	return len(a)
}

func (a ByLength) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByLength) Less(i, j int) bool {
	return (a[i] + a[j]) > (a[j] + a[i])
}

func largestNumber(nums []int) string {
	if nums != nil {
		if len(nums) == 0 {
			return ""
		} else if len(nums) == 1 {
			return strconv.Itoa(nums[0])
		}
	}else{
		return ""
	}
	l := make([]string, 0)

	for _, v := range nums {
		l = append(l, strconv.Itoa(v))
	}
	sort.Sort(ByLength(l))

	numstr := strings.Join(l, "")
	numstr = strings.TrimLeft(numstr, "0")
	if numstr =="" {
		return "0"
	}
	return numstr
}

func main() {
	nums := []int{0,-1}
	fmt.Println(largestNumber(nums))
	nums = []int{10, 2}
	fmt.Println(largestNumber(nums))
}
