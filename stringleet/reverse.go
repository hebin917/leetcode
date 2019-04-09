package main

import (
	"fmt"
	"strings"
	"bytes"
)

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(s[:])
	fmt.Println(reverseWords(s))
	fmt.Println(reverseString(s))
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}


func reverseString1(s string) string {
	res := ""
	for _, v := range s {
		res =string(v)+res
	}
	return res
}


func reverseString(s string) string {
	var res=bytes.Buffer{}
	for i:=len(s) -1;i>=0;i-- {
		res.WriteByte(s[i])
	}

	return res.String()
}


func reverseWords(s string) string {

	tmpStr := strings.Split(s, " ")
	res := make([]string, len(tmpStr))
	var i int = 0
	if s != "" {
		for _, str := range tmpStr {

			res[i] = Reverse(str)
			i++
		}
	}
	return strings.Join(res, " ")
}


