package main

import "fmt"

/*
890. Find and Replace Pattern
You have a list of words and a pattern, and you want to know which words in words matches the pattern.

A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.

(Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.)

Return a list of the words in words that match the given pattern.

You may return the answer in any order.



Example 1:

Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
Output: ["mee","aqq"]
Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
"ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
since a and b map to the same letter.

 */

 /*
 如果模式匹配的话，所有字母映射的的数字是相等的
  */
func findAndReplacePattern(words []string, pattern string) []string {
	res := []string{}
	for _, w := range words {
		p := make([]int, 26)
		s := make([]int, 26)
		same := true
		for i, v := range w {
			if s[v-'a'] != p[pattern[i]-'a'] {
				same = false
				break
			} else {
				x:=i+1   //不能为i，因为默认为0，
				s[v-'a'] = x
				p[pattern[i]-'a'] = x
			}
		}
		if same {
			res=append(res,w)
		}
	}
	return res
}

func main() {
	words:=[]string {"abc","deq","mee","aqq","dkd","ccc"}
	pattern:="abb"
	res := findAndReplacePattern(words,pattern)
	for _,v :=range res {
		fmt.Println(v)
	}
}
