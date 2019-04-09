package main

import "fmt"

/*

852. Peak Index in a Mountain Array

Let's call an array A a mountain if the following properties hold:

A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:

Input: [0,1,0]
Output: 1
Example 2:

Input: [0,2,1,0]
Output: 1


 */
func peakIndexInMountainArray(A []int) int {
	if len(A) <3 {
		return -1
	}
	l:=1
	r:=len(A)-1
	var mid int
	for l<r {
		mid = (l+r)/2
		if A[mid] <A[mid+1] {
			l=mid
		}else if A[mid-1] > A[mid] {
			r=mid
		}else {
			return mid
		}
	}
	return -1
}


func main() {
	A:= []int {0,1,0}
	fmt.Println(peakIndexInMountainArray(A))
	A=[]int {0,2,1,0}
	fmt.Println(peakIndexInMountainArray(A))
}