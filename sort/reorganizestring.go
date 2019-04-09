package main

import (
	"strings"
	"fmt"
)

/*
767
Given a string S, check if the letters can be rearranged so that two characters that are adjacent to each other are not the same.

If possible, output any possible result.  If not possible, return the empty string.

Example 1:

Input: S = "aab"
Output: "aba"
Example 2:

Input: S = "aaab"
Output: ""

 */


 /*
 解题思想： 一个数组的序列除了奇数就是偶数，利用奇偶来进行分别插入：
 1、首先求得数组中每个元素的数据
 2、求得个数最大的元素的index
 3、将最大元素放入res中
 4、在放入其它元素，如果其它index+2>=len,则将index=1
 如果最大元素的个数*2 -1 > len(S)的话，表示不成立
  */
func reorganizeString(S string) string {
	ch := make([]int, 26)
	max := 0
	for _, c := range S {
		ch[c-'a'] ++
		if ch[c-'a'] > ch[max] {
			max = int(c - 'a')
		}
	}

	l := len(S)
	if l < 2*ch[max] -1 {
		return ""
	}

	idx := 0
	res := make([]string, len(S))
	for i:=0; i< ch[max];i++ {
		res[idx]=string(max+'a')
		idx+=2
	}
	ch[max]=0
	for c, v := range ch {
		if v == 0 {
			continue
		}
		for i := 0; i < v; i++ {
			if idx >=len(S) {
				idx = 1
			}
			res[idx] = string(c + 'a')
			idx += 2
		}
	}

	return strings.Join(res, "")
}

func main() {
	S := "aabbbbbbbbbbbaaaaaddddaaaarrrccccddessseaaaaaaa"
	fmt.Println(reorganizeString(S))
	fmt.Println("----")

	S = "aab"
	fmt.Println(reorganizeString(S))
	fmt.Println("----")
	S = "aabb"
	fmt.Println(reorganizeString(S))
	fmt.Println("----")

}
