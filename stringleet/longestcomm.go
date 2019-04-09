package main

import (
	"strings"
	"fmt"
)

func main() {
	strs := []string {"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string {"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) ==0 {
		return ""
	}
	minstr := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(minstr) {
			minstr = strs[i]
		}
	}

	j := 0
	for (j < len(strs)) {
		for strings.Index(strs[j], minstr) != 0 {
			minstr = minstr[0 : len(minstr)-1]
		}
		j++
	}
	return minstr

}
