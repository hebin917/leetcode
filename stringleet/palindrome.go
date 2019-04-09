package main

import "fmt"

func main() {
	s:="abdcba"
	start:= len(s)-1
	fmt.Println(s[start:1])
	fmt.Println(validPalindrome(s))
}

func validPalindrome(s string) bool {
	i := -1
	r := len(s)
	for  {
		i++
		r--
		if i < r {
			if s[i] != s[r] {
				return isPalindromic(s, i, r+1) || isPalindromic(s, i-1, r)
			}
		}else {
			return true
		}
	}
}

func isPalindromic(s string,st int,e int) bool {
	for {
		st++
		e--
		if st < e {
			if s[st] != s[e] {
				return false
			}
		}else {
			break
		}
	}
	return true
}
