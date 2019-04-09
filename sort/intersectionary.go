package main

import "fmt"

/**
chacha'zh查找俩个nums的集合数据
 */
func intersect(nums1 []int,nums2 []int)[]int{
	m:=make(map[int]int)
	res :=[]int{}
	for _,v :=range nums1 {
		m[v]++
	}
	for _,v :=range nums2 {
		if _,ok := m[v];ok{
			m[v]--
			if m[v] >=0 {
				res=append(res,v)
			}
		}
	}
	return res
}


func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersect(nums1,nums2))
	fmt.Println(intersect(nums2,nums1))
	nums1 = [] int {4,9,5}
	nums2 = [] int {9,4,9,8,4}
	fmt.Println(intersect(nums1,nums2))
	fmt.Println(intersect(nums2,nums1))
}