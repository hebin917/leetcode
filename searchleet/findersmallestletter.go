package main

import (
	"fmt"
)

/*
744. Find Smallest Letter Greater Than Target

Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.

Letters also wrap around. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.
 */

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)

	lo, hi := 0, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if letters[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	//如果在index=n时结束，这时就需要返回第一个elements
	return letters[lo %n]
}

func main() {
	letters := []byte{'c', 'f', 'j'}
	fmt.Println(int(nextGreatestLetter(letters, 'a')))
	letters=[]byte{'c', 'f', 'j'}
	fmt.Println(int(nextGreatestLetter(letters, 'k')))

}
