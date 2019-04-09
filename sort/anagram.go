package main

import "fmt"

/*
猜字谜： 给定俩个字符串，然后判断字符串中是否包含所有的字母。

 */
func isAnagram(s string, t string) bool {
	smap := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		smap[s[i]] += 1
	}
	for i := 0; i < len(t); i++ {
		smap[t[i]] --
	}

	for _, v := range smap {
		if v !=0 {
			return false
		}
	}

	return true
}

func isAnagram1(s string, t string) bool {
	res := make([]uint8, 26)
	for i := 0; i < len(s); i++ {
		res[s[i]-'a']++
	}
	for j := 0; j < len(t); j++ {
		res[t[j]-'a']--
	}
	fmt.Println(res)
	for _, i := range res {
		if i != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
	s = "rat"
	t = "car"
	fmt.Println(isAnagram(s, t))
}
